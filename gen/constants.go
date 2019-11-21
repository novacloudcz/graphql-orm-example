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

enum ObjectSortType {
  ASC
  DESC
}

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

input CompanySortType {
  id: ObjectSortType
  name: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  employeesIds: ObjectSortType
  employees: UserSortType
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
  id_null: Boolean
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
  name_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
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

input UserSortType {
  id: ObjectSortType
  email: ObjectSortType
  firstName: ObjectSortType
  lastName: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  tasksIds: ObjectSortType
  companiesIds: ObjectSortType
  friendsIds: ObjectSortType
  tasks: TaskSortType
  companies: CompanySortType
  friends: UserSortType
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
  id_null: Boolean
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
  email_null: Boolean
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
  firstName_null: Boolean
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
  lastName_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
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

input TaskSortType {
  id: ObjectSortType
  title: ObjectSortType
  completed: ObjectSortType
  dueDate: ObjectSortType
  type: ObjectSortType
  description: ObjectSortType
  assigneeId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  assignee: UserSortType
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
  id_null: Boolean
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
  title_null: Boolean
  completed: Boolean
  completed_ne: Boolean
  completed_gt: Boolean
  completed_lt: Boolean
  completed_gte: Boolean
  completed_lte: Boolean
  completed_in: [Boolean!]
  completed_null: Boolean
  dueDate: Time
  dueDate_ne: Time
  dueDate_gt: Time
  dueDate_lt: Time
  dueDate_gte: Time
  dueDate_lte: Time
  dueDate_in: [Time!]
  dueDate_null: Boolean
  type: TaskType
  type_ne: TaskType
  type_gt: TaskType
  type_lt: TaskType
  type_gte: TaskType
  type_lte: TaskType
  type_in: [TaskType!]
  type_null: Boolean
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
  description_null: Boolean
  assigneeId: ID
  assigneeId_ne: ID
  assigneeId_gt: ID
  assigneeId_lt: ID
  assigneeId_gte: ID
  assigneeId_lte: ID
  assigneeId_in: [ID!]
  assigneeId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  assignee: UserFilterType
}

type TaskResultType {
  items: [Task!]!
  count: Int!
}`
)
