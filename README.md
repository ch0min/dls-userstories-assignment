# Project: User Stories Assignment

## Table of Contents

- [Project: User Stories Assignment](#project-user-stories-assignment)
  - [Table of Contents](#table-of-contents)
  - [Project Description](#project-description)
  - [Project Board Usage](#project-board-usage)
    - [Issue Management Procedure](#issue-management-procedure)
    - [Labels](#labels)
    - [Estimation](#estimation)
    - [Issue Lifecycle](#issue-lifecycle)
    - [Responsibilities](#responsibilities)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [After Thought](#after-thought)

## Project Description

This project aims to document and develop user stories for the Chemical Warehouse Assignment. The application will fulfill the user story's functional and non-functional requirements. The project utilizes Golang as the backend.

---

## Project Board Usage

### Issue Management Procedure

The team will manage tasks and features through a structured process using the project board. Below is a breakdown of the procedure:

### Labels

Each issue will be tagged with relevant labels for easy categorization:

- **help wanted**: If you are stuck and need help.
- **question**: If you have a question or clarification about a task.
- **documentation**: For changes or additions to project documentation.
- **priority**: To indicate the urgency of the task we used the following labels:
  - `P0` (Critical Priority)
  - `P1` (High Priority)
  - `P2` (Medium Priority)

### Estimation

Each issue will have a time estimate assigned, helping the team gauge how long tasks are expected to take. Estimates will be recorded in hours or days and labeled accordingly.

- Example:
  - `estimate: 2h`
  - `estimate: 1d`

### Issue Lifecycle

1. **Creation**: Team members can create issues by clearly describing the problem or feature request, adding relevant labels, and assigning an estimated time.
2. **Assignment**: Issues will be assigned to team members based on availability and expertise.
3. **Progress**: Issues will be moved through the board as they progress (`Backlog` → `In Progress` → `Review` → `Testing` → `Done`).
4. **Completion**: Upon resolving the issue, it will be marked as `Done`, and the pull request will be reviewed by another team member.

### Responsibilities

- **Issue Assignment**: The project lead will assign issues during sprint planning as necessary.
- **Team Collaboration**: Each team member is responsible for regularly updating the status of the issues they are assigned and keeping track of time estimates.

---

## Getting Started

### Prerequisites

- go mod tidy

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ch0min/dls-userstories-assignment
   ```
2. start docker desktop

3. run `docker-compose up -d` to start the database (make sure that you are in the /server directory)

4. run `go run main.go` to start the server (make sure that you are in the root directory)

### After Thought

- When our project scales and become more exam focused, we'll make more labels, worksflow etc. There will be a more detailed display.
