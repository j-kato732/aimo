CREATE TABLE "period_models"  (
  "id" INTEGER NOT NULL PRIMARY KEY,
  "period" TEXT
)

CREATE TABLE "job_models"  (
  "id" INTEGER NOT NULL PRIMARY KEY,
  "job" TEXT
)

CREATE TABLE `policy_models`  (
  `created_at` datetime,
  `id` integer,
  `mid_term_policy` text,
  `period` text,
  `period_policy` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
)

CREATE TABLE `role_models`  (
  `created_at` datetime,
  `department_id` integer,
  `id` integer,
  `job_id` integer,
  `period` text,
  `role` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
)

CREATE TABLE `user_models`  (
  `admin_flg` numeric,
  `auth_id` integer,
  `created_at` datetime,
  `department_id` integer,
  `enrollment_flg` numeric,
  `first_name` text,
  `id` integer,
  `job_id` integer,
  `last_name` text,
  `period` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
)

CREATE TABLE `achievement_mean_models`  (
  `achievement_mean` text,
  `achievement_mean_number` integer,
  `aim_number` integer,
  `created_at` datetime,
  `fifth_month` numeric,
  `first_month` numeric,
  `fourth_month` numeric,
  `id` integer,
  `period` text,
  `second_month` numeric,
  `sixth_month` numeric,
  `third_month` numeric,
  `updated_at` datetime,
  `user_id` integer,
  PRIMARY KEY (`id`)
)

CREATE TABLE `aim_models`  (
  `achievement_difficulty_before` integer,
  `achievement_level` text,
  `achievement_weight` integer,
  `aim_item` text,
  `aim_number` integer,
  `created_at` datetime,
  `id` integer,
  `period` text,
  `updated_at` datetime,
  `user_id` integer,
  PRIMARY KEY (`id`)
)

CREATE TABLE `comprehensive_comment_models`  (
  `comment_user_id` integer,
  `comprehensive_comment` text,
  `id` integer,
  `period` text,
  `user_id` integer,
  PRIMARY KEY (`id`)
)

CREATE TABLE `department_goal_models`  (
  `created_at` datetime,
  `department_goal` text,
  `department_id` integer,
  `id` integer,
  `period` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
)

CREATE TABLE "department_models"  (
  "id" INTEGER NOT NULL PRIMARY KEY,
  "department" TEXT
)

CREATE TABLE `evaluation_before_models`  (
  `aim_id` integer,
  `comment` text,
  `comment_user_id` integer,
  `evaluator_number` integer,
  `id` integer,
  PRIMARY KEY (`id`)
)

CREATE TABLE `evaluation_models`  (
  `achievement_difficulty` integer,
  `achievement_rate` integer,
  `aim_id` integer,
  `comment` text,
  `evaluator_number` integer,
  `evaluator_user_id` integer,
  `id` integer,
  PRIMARY KEY (`id`)
)

CREATE TABLE `personal_eva_models`  (
  `achievement_difficulty` integer,
  `achievement_rate` integer,
  `aim_id` integer,
  `id` integer,
  `personal_evaluation` text,
  PRIMARY KEY (`id`)
)

INSERT INTO period_models (period) VALUES (202105);
INSERT INTO period_models (period) VALUES (202111);