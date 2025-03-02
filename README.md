# Task Manager CLI 📋

[![Go Version](https://img.shields.io/badge/go-1.21%2B-blue.svg)](https://golang.org/doc/install)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://makeapullrequest.com)

A feature-rich command-line task manager with persistent storage and color-coded interface. Manage your tasks efficiently with this intuitive Go-based CLI tool. This project follows the [Task Tracker Project Roadmap](https://roadmap.sh/projects/task-tracker) from roadmap.sh.

![CLI Demo](demo.gif) <!-- Add actual demo gif later -->

## Table of Contents
- [Project Roadmap](#project-roadmap-)
- [Features](#features-)
- [Installation](#installation-)
- [Usage](#usage-)
- [Examples](#examples-)
- [Color Coding](#color-coding-)
- [Data Storage](#data-storage-)
- [Contributing](#contributing-)
- [License](#license-)

## Project Roadmap 🗺️
This implementation follows the recommended development path from [roadmap.sh's Task Tracker Project Guide](https://roadmap.sh/projects/task-tracker). The roadmap helped shape:

- Core feature selection
- Architecture decisions
- Implementation priorities
- Best practice adoption
- Error handling strategies
- Persistence layer design

## Features ✨

- **CRUD Operations**: Full Create, Read, Update, Delete functionality
- **Status Tracking**: Three states - `pending`, `in-progress`, and `done`
- **Colorful UI**: Visual status indicators with ANSI colors
- **Persistent Storage**: Automatic saving to JSON file
- **Filtering**: List tasks by specific status
- **Bulk Operations**: Clear all tasks with one command
- **Help System**: Built-in documentation with examples
- **Error Handling**: Friendly error messages with troubleshooting tips
- **ID Management**: Automatic ID generation and maintenance

## Installation ⚙️

### Prerequisites
- Go 1.21 or newer
- Git (for source installation)

### Methods

#### Using Go Install
```bash
go install github.com/dangbros/task-tracker@latest