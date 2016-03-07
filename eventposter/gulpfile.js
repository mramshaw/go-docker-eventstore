var gulp = require('gulp');
var child = require('child_process');
var notifier   = require('node-notifier');
var util       = require('gulp-util');
var server = null;

//Runs the file watching and then calls init
gulp.task('default', ['watch', 'init']);

//Runs the initial build and run
gulp.task('init', ['build', 'run']);

//Sets up file watching
gulp.task('watch', function() {
  gulp.watch('./**/*.go', ['build','run']);
  gulp.watch('./**/*.html', ['build','run']);
});

//Call go install which will compile the go code
//and install it in the GOBIN directory
gulp.task('build', function() {
  var build = child.spawnSync('go', ['install']);
  if (build.stderr.length) {
    var lines = build.stderr.toString()
      .split('\n').filter(function(line) {
        return line.length
      });
    for (var l in lines)
      util.log(util.colors.red(
        'Error (go install): ' + lines[l]
      ));
    notifier.notify({
      title: 'Error (go install)',
      message: lines
    });
  }
  return build;
});

//Stop the running process and run a new one
gulp.task('run', function() {
  if (server)
    server.kill();
  server = child.spawn('app');
  server.stderr.on('data', function(data) {
    process.stdout.write(data.toString());
  });
  server.stdout.on('data', function(data) {
    process.stdout.write(data.toString());
  });
});