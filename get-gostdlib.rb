#!/usr/bin/env ruby

$pkg_ignores = [
  'hash',
  'hash/crc64',
  'encoding'
]

$symbol_ignores = [
  ['math', 'MaxUint64']
]

require 'set'

class Package
  attr_reader :name, :pname, :info

  def initialize(name, pname, info)
    @name = name
    @pname = pname
    @ref = File.basename(@name)
    @info = info
    @types = Set.new()
    @functions = Set.new()
    @constants = Set.new()
    @variables = Set.new()
    @info.each do |i|
      if m = /^type (.+?) struct.+$/.match(i)
        @types.add m[1] unless $symbol_ignores.include?([name, m[1]])
      end
      if m = /^func (.+?)\(.*$/.match(i)
        @functions.add m[1] unless $symbol_ignores.include?([name, m[1]])
      end
      if m = /^const (.+?) =.*$/.match(i)
        @constants.add m[1] unless $symbol_ignores.include?([name, m[1]])
      end
      if m = /^var (.+?) .*$/.match(i)
        @variables.add m[1] unless $symbol_ignores.include?([name, m[1]])
      end
    end
  end

  def write
    begin
      Dir.mkdir('gostdlib')
    rescue => e
    end
    begin
      Dir.mkdir(File.join('gostdlib', @pname))
    rescue => e
    end
    s = 
      "package main\n"\
      "\n"\
      "import (\n"\
      "  lua \"github.com/yuin/gopher-lua\"\n"\
      "  luar \"layeh.com/gopher-luar\"\n"\
      "  \"#{@name}\"\n"\
      ")\n"\
      "\n"\
      "func GoglingLoad(L *lua.LState) {\n"\
      "  L.PreloadModule(\"go.#{@pname}\", loader)\n"\
      "}\n"\
      "\n"\
      "func loader(L *lua.LState) int {\n"\
      "  tbl := L.NewTable()\n"\
      "  L.SetField(tbl, \"_go_pkg\", lua.LString(\"#{name}\"))\n"\
      "  // Types\n"
    @types.each do |t|
      s << "  L.SetField(tbl, \"#{t}\", luar.New(L, func() *#{@ref}.#{t} { return &#{@ref}.#{t}{} }))\n"
    end
    s << "  // Functions\n"
    @functions.each do |f|
      s << "  L.SetField(tbl, \"#{f}\", luar.New(L, #{@ref}.#{f}))\n"
    end
    s << "  // Constants\n"
    @constants.each do |c|
      s << "  L.SetField(tbl, \"#{c}\", luar.New(L, #{@ref}.#{c}))\n"
    end
    s << "  // Variables\n"
    @variables.each do |v|
      s << "  L.SetField(tbl, \"#{v}\", luar.New(L, &#{@ref}.#{v}))\n"
    end
    s << "  // end\n  L.Push(tbl)\n  return 1\n}"
    File.write(File.join('gostdlib', @pname, "#{@pname}.go"), s)
  end

  def to_s
    "<Package name=#{@name} pname=#{@pname}>"
  end

  def inspect
    to_s
  end
end

$packages = {}
$package_names = Set.new
$package_info = {}

$goroot = ENV['GOROOT']

$apidir = File.join($goroot, 'api')

$apifiles = File.join($apidir, 'go1*.txt')
$apifiles = Dir[$apifiles]
$apifiles << File.join($apidir, 'go1.txt')

$lines = []

$apifiles.each do |fname|
  f = File.open(fname, 'r')
  while s = f.gets
    $lines << s
  end
  f.close
end

def add_package(pkg, info)
  return if $pkg_ignores.include?(pkg)
  $package_names.add pkg
  unless $package_info[pkg]
    $package_info[pkg] = []
  end
  $package_info[pkg] << info
end

def process(name)
  name.gsub(/\/(.)/) { |s| s[1].upcase }
end

$lines.each do |l|
  m = /^pkg ([^()\s]*?) \(linux-amd64\), (.*)$/.match l
  if m
    add_package(m[1], m[2])
  else
    m = /^pkg ([^()\s]*?), (.*)$/.match l
    unless m
      raise "invalid line: #{l}" unless /^#.*/.match?(l) or /^\s*$/.match?(l) or /^pkg ([^()\s]*?) \(.*?\), (.*)$/.match?(l)
    else
      add_package(m[1], m[2])
    end
  end
end

$package_names.each do |name|
  pname = process(name)
  $packages[pname] = Package.new(name, pname, $package_info[name])
end

$packages.each_pair do |k, v|
  puts k, v
  v.write
end

File.write("packages.sh", "export GO2LUAPKGS=#{$package_names.to_a.map{|s|process(s)}.join(':')}")