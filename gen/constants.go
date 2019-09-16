package gen

type key int

const (
	KeyPrincipalID      key    = iota
	KeyLoaders          key    = iota
	KeyExecutableSchema key    = iota
	KeyJWTClaims        key    = iota
	SchemaSDL           string = `scalar Time

type Query {
  company(id: ID, q: String, filter: CompanyFilterType): Company
  companies(offset: Int, limit: Int = 30, q: String, sort: [CompanySortType!], filter: CompanyFilterType): CompanyResultType
  user(id: ID, q: String, filter: UserFilterType): User
  users(offset: Int, limit: Int = 30, q: String, sort: [UserSortType!], filter: UserFilterType): UserResultType
  task(id: ID, q: String, filter: TaskFilterType): Task
  tasks(offset: Int, limit: Int = 30, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType
}

type Mutation {
  createCompany(input: CompanyCreateInput!): Company!
  updateCompany(id: ID!, input: CompanyUpdateInput!): Company!
  deleteCompany(id: ID!): Company!
  deleteAllCompanies: Boolean!
  createUser(input: UserCreateInput!): User!
  updateUser(id: ID!, input: UserUpdateInput!): User!
  deleteUser(id: ID!): User!
  deleteAllUsers: Boolean!
  createTask(input: TaskCreateInput!): Task!
  updateTask(id: ID!, input: TaskUpdateInput!): Task!
  deleteTask(id: ID!): Task!
  deleteAllTasks: Boolean!
}

directive @validator(required: Boolean!) on FIELD_DEFINITION

extend type Query {
  hello: String!
}

type Company {
  id: ID!
  name: String
  employees: [User!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  employeesIds: [ID!]!
}

type User {
  id: ID!
  email: String
  firstName: String
  lastName: String
  tasks: [Task!]!
  companies: [Company!]!
  friends: [User!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  tasksIds: [ID!]!
  companiesIds: [ID!]!
  friendsIds: [ID!]!
}

enum TaskType {
  BUG
  TASK
}

type Task {
  id: ID!
  title: String
  completed: Boolean
  dueDate: Time
  type: TaskType
  description: String
  assignee: User
  assigneeId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

input CompanyCreateInput {
  id: ID
  name: String
  employeesIds: [ID!]
}

input CompanyUpdateInput {
  name: String
  employeesIds: [ID!]
}

enum CompanySortType {
  ID_ASC
  ID_DESC
  NAME_ASC
  NAME_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  EMPLOYEES_IDS_ASC
  EMPLOYEES_IDS_DESC
}

input CompanyFilterType {
  AND: [CompanyFilterType!]
  OR: [CompanyFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  employees: UserFilterType
}

type CompanyResultType {
  items: [Company!]!
  count: Int!
}

input UserCreateInput {
  id: ID
  email: String
  firstName: String
  lastName: String
  tasksIds: [ID!]
  companiesIds: [ID!]
  friendsIds: [ID!]
}

input UserUpdateInput {
  email: String
  firstName: String
  lastName: String
  tasksIds: [ID!]
  companiesIds: [ID!]
  friendsIds: [ID!]
}

enum UserSortType {
  ID_ASC
  ID_DESC
  EMAIL_ASC
  EMAIL_DESC
  FIRST_NAME_ASC
  FIRST_NAME_DESC
  LAST_NAME_ASC
  LAST_NAME_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  TASKS_IDS_ASC
  TASKS_IDS_DESC
  COMPANIES_IDS_ASC
  COMPANIES_IDS_DESC
  FRIENDS_IDS_ASC
  FRIENDS_IDS_DESC
}

input UserFilterType {
  AND: [UserFilterType!]
  OR: [UserFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  email: String
  email_ne: String
  email_gt: String
  email_lt: String
  email_gte: String
  email_lte: String
  email_in: [String!]
  email_like: String
  email_prefix: String
  email_suffix: String
  firstName: String
  firstName_ne: String
  firstName_gt: String
  firstName_lt: String
  firstName_gte: String
  firstName_lte: String
  firstName_in: [String!]
  firstName_like: String
  firstName_prefix: String
  firstName_suffix: String
  lastName: String
  lastName_ne: String
  lastName_gt: String
  lastName_lt: String
  lastName_gte: String
  lastName_lte: String
  lastName_in: [String!]
  lastName_like: String
  lastName_prefix: String
  lastName_suffix: String
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  tasks: TaskFilterType
  companies: CompanyFilterType
  friends: UserFilterType
}

type UserResultType {
  items: [User!]!
  count: Int!
}

input TaskCreateInput {
  id: ID
  title: String
  completed: Boolean
  dueDate: Time
  type: TaskType
  description: String
  assigneeId: ID
}

input TaskUpdateInput {
  title: String
  completed: Boolean
  dueDate: Time
  type: TaskType
  description: String
  assigneeId: ID
}

enum TaskSortType {
  ID_ASC
  ID_DESC
  TITLE_ASC
  TITLE_DESC
  COMPLETED_ASC
  COMPLETED_DESC
  DUE_DATE_ASC
  DUE_DATE_DESC
  TYPE_ASC
  TYPE_DESC
  DESCRIPTION_ASC
  DESCRIPTION_DESC
  ASSIGNEE_ID_ASC
  ASSIGNEE_ID_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
}

input TaskFilterType {
  AND: [TaskFilterType!]
  OR: [TaskFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  title: String
  title_ne: String
  title_gt: String
  title_lt: String
  title_gte: String
  title_lte: String
  title_in: [String!]
  title_like: String
  title_prefix: String
  title_suffix: String
  completed: Boolean
  completed_ne: Boolean
  completed_gt: Boolean
  completed_lt: Boolean
  completed_gte: Boolean
  completed_lte: Boolean
  completed_in: [Boolean!]
  dueDate: Time
  dueDate_ne: Time
  dueDate_gt: Time
  dueDate_lt: Time
  dueDate_gte: Time
  dueDate_lte: Time
  dueDate_in: [Time!]
  type: TaskType
  type_ne: TaskType
  type_gt: TaskType
  type_lt: TaskType
  type_gte: TaskType
  type_lte: TaskType
  type_in: [TaskType!]
  description: String
  description_ne: String
  description_gt: String
  description_lt: String
  description_gte: String
  description_lte: String
  description_in: [String!]
  description_like: String
  description_prefix: String
  description_suffix: String
  assigneeId: ID
  assigneeId_ne: ID
  assigneeId_gt: ID
  assigneeId_lt: ID
  assigneeId_gte: ID
  assigneeId_lte: ID
  assigneeId_in: [ID!]
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  assignee: UserFilterType
}

type TaskResultType {
  items: [Task!]!
  count: Int!
}`
)