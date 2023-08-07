task :is_repo_clean do
    abort 'please commit your changes first!' unless `git status -s | wc -l`.strip.to_i.zero?
  end
  
  task :current_branch do
    `git rev-parse --abbrev-ref HEAD`.strip
  end
  
  task :command_exists, [:command] do |_, args|
    abort "#{args.command} doesn't exists" if `command -v #{args.command} > /dev/null 2>&1 && echo $?`.chomp.empty?
  end
  
  task :has_bumpversion do
      Rake::Task['command_exists'].invoke('bumpversion')
  end
  
  task :has_godoc do
      Rake::Task['command_exists'].invoke('godoc')
  end
  
  task :golangcilint do
    Rake::Task['command_exists'].invoke('golangci-lint')
  end
  
  def current_version(lookup_file=".bumpversion.cfg")
    file = File.open(lookup_file, "r")
    data = file.read
    file.close
    match = /current_version = (\d+).(\d+).(\d+)/.match(data)
    "#{match[1]}.#{match[2]}.#{match[3]}"
  end
     
  desc "run tests"
  task :test do
    system "go test -p 1 -v -race ./..."
  end
  
  AVAILABLE_REVISIONS = %w[major minor patch].freeze
  desc "bump version, default is: patch"
  task :bump, [:revision] => [:has_bumpversion] do |_, args|
    args.with_defaults(revision: 'patch')
    unless AVAILABLE_REVISIONS.include?(args.revision)
      abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}"
    end
  
    system "bumpversion #{args.revision}"
  end
  
  PACKAGE_NAME = "github.com/erhanakp/sugaredgorm"
  desc 'run doc server'
  task :doc, [:port] => [:has_godoc] do |_, args|
    args.with_defaults(port: 9009)
  
    puts "running doc server at \e[33m:#{args.port}\e[0m\n"
    puts "\t\e[33mhttp://localhost:#{args.port}/pkg/#{PACKAGE_NAME}/\e[0m"
    puts "\n"
  
    system "godoc -http=:#{args.port}"
  end
  
  desc "publish new version of the library, default is: patch"
  task :publish, [:revision] => [:is_repo_clean] do |t, args|
    current_branch = Rake::Task["current_branch"].invoke.first.call
  
    abort "this command works for [main] or [master] branches only, not for [#{current_branch}] branch" unless ['master', 'main'].include?(current_branch)
  
    args.with_defaults(revision: "patch")
  
    Rake::Task["bump"].invoke(args.revision)
    current_git_tag = "v#{current_version}"
  
    puts "[->] new version is \e[33m#{current_git_tag}\e[0m"
    puts "[->] pushing \e[33m#{current_git_tag}\e[0m to remote"
    system %{
      git push origin #{current_git_tag} &&
      go list -m github.com/erhanakp/sugaredgorm@#{current_git_tag} &&
      echo "[->] [#{current_git_tag}] has been published" &&
      git push origin #{current_branch} &&
      echo "[->] code pushed to: [#{current_branch}] branch (updated)"
    }
  end
  
  desc "run golangci-lint"
  task :lint => [:golangcilint] do
    system "golangci-lint run"
  end