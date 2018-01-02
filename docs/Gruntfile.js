var grunt = require('grunt');
grunt.loadNpmTasks('grunt-apidoc');
 grunt.initConfig({
	apidoc: {
    myapp: {
      src: "app/",
      dest: "apidoc/"
    }
  }
 });
 
 grunt.registerTask('default', ['apidoc']); 