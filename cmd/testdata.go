package cmd

var (
	TestYmlStr = `
notes:
 goal:
   - to test using a simpler mode of data structure for tasks using array
 why:
   - array will allow you to put more attribute so that you can put desc

tasks:
 -
   name: task1
   desc: this is task1
   task: #this comment will not be treated as desc of the task, removing # will invalid the yml
     -
       func: shell
       desc: do step1 in shell func
       do:
         - echo "hello"
         - echo "world"

     -
       func: cmd
       desc: do step2 in shell func
       flags:
         - ignore_error
       do:
         - echo "hello"
         - echo "world"|grep non-exist
 -
   name: task2
   desc: this is task2
   task: #this comment will not be treated as desc of the task, removing # will invalid the yml
     -
       func: shell
       desc: do step1 in shell func
       do:
         - echo "hello"
         - echo "world"

     -
       func: cmd
       desc: do step2 in shell func
       flags:
         - ignore_error
       do:
         - echo "hello"
         - echo "world"|grep non-exist

`

	TestYmlStr2 = `
name: Fred
age: 22
---
name: Stella
age: 23
---
name: Android
age: 232
`

	TestYmlStr3 = `
nsw:
  sydney:
    -
      sg:
        student:
          name: Grace
          gender: Female
          school: MLC
    -
      kings:
        student:
          name: Emily
          gender: Female
          school: KINGS
`
)
