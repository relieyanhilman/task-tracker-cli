# Task Tracker CLI

## Introduction

Task Tracker CLI is a command line application designed to help you track and manage your tasks. With this application, you can monitor what you need to do, what you've done, and what you're currently working on.

## Technologies

This application is built using Go Programming Language and Go Internal Library.

## Launch

To run this application, you can follow these steps:

1. Build the project by running the command `go build -o task-cli`
2. Run the project by running the command `./task-cli <command>` (on Windows)

## Table of Contents

* [Introduction](#introduction)
* [Technologies](#technologies)
* [Launch](#launch)
* [Scope Functionalities](#scope-functionalities)
* [Examples of Use](#examples-of-use)
* [Project Status](#project-status)
* [Sources](#sources)
* [Project Idea URL](#project-idea-url)

## Scope Functionalities

This application has several commands that can be used to manage your tasks. Here is a list of available commands:

* `add`: Add a new task
* `update`: Update an existing task
* `delete`: Delete an existing task
* `list`: Display a list of existing tasks
* `list todo`: Display a list of existing tasks with the status of "to do"
* `list in-progress`: Display a list of existing tasks with the status of "in progress"
* `list done`: Display a list of existing tasks with the status of "done"
* `mark-in-progress`: Mark a task as in progress
* `mark-done`: Mark a task as done

## Examples of Use

Here are examples of how to use some of the available commands:

* Add a new task: `./task-cli add "Create a report"`
* Update an existing task: `./task-cli update 1 "Create a better report"`
* Delete an existing task: `./task-cli delete 1`
* Display a list of all existing tasks: `./task-cli list`
* Display a list of existing tasks with the status of "to do": `./task-cli list todo`
* Display a list of existing tasks with the status of "in progress": `./task-cli list in-progress`
* Display a list of existing tasks with the status of "done": `./task-cli list done`
* Mark a task as in progress: `./task-cli mark-in-progress 1`
* Mark a task as done: `./task-cli mark-done 1`

## Project Status

This application is complete and ready to use.

## Sources

This application uses Go Documentation as a reference source.

## Project Idea URL

The idea for this project originated from the following link: https://roadmap.sh/projects/task-tracker
