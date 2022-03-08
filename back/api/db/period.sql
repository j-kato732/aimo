CREATE TABLE "period_models"  (
  "id" INTEGER NOT NULL PRIMARY KEY,
  "period" TEXT
);

CREATE TABLE "job_models"  (
  "id" INTEGER NOT NULL PRIMARY KEY,
  "job" TEXT
);

CREATE TABLE `policy_models`  (
  `created_at` datetime,
  `id` integer,
  `mid_term_policy` text,
  `period` text,
  `period_policy` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
);

CREATE TABLE `role_models`  (
  `created_at` datetime,
  `department_id` integer,
  `id` integer,
  `job_id` integer,
  `period` text,
  `role` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
);

CREATE TABLE `user_models`  (
  `admin_flg` numeric,
  `auth_id` text,
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
);

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
);

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
);

CREATE TABLE `comprehensive_comment_models`  (
  `comment_user_id` integer,
  `comprehensive_comment` text,
  `id` integer,
  `period` text,
  `user_id` integer,
  PRIMARY KEY (`id`)
);

CREATE TABLE `department_goal_models`  (
  `created_at` datetime,
  `department_goal` text,
  `department_id` integer,
  `id` integer,
  `period` text,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
);

CREATE TABLE "department_models"  (
  "id" INTEGER NOT NULL PRIMARY KEY,
  "department" TEXT
);

CREATE TABLE `evaluation_before_models`  (
  `aim_id` integer,
  `comment` text,
  `comment_user_id` integer,
  `evaluator_number` integer,
  `id` integer,
  PRIMARY KEY (`id`)
);

CREATE TABLE `evaluation_models`  (
  `achievement_difficulty` integer,
  `achievement_rate` integer,
  `aim_id` integer,
  `comment` text,
  `evaluator_number` integer,
  `evaluator_user_id` integer,
  `id` integer,
  PRIMARY KEY (`id`)
);

CREATE TABLE `personal_eva_models`  (
  `achievement_difficulty` integer,
  `achievement_rate` integer,
  `aim_id` integer,
  `id` integer,
  `personal_evaluation` text,
  PRIMARY KEY (`id`)
);

INSERT INTO period_models (period) VALUES (202105);
INSERT INTO period_models (period) VALUES (202111);

INSERT INTO department_models (department) VALUES ("ソリューション本部 開発1部"),
("ソリューション本部 開発2部"),
("エイジア開発事業本部 オフショア/エボ開発部"),
("エイジア開発事業本部 UXデザイン開発部"),
("エイジア開発事業本部 デジタルソリューション部"),
("札幌開発センター 開発部"),
("札幌開発センター 営業部"),
("営業本部 営業1G"),
("営業本部 営業2G"),
("事業推進本部 リクルーティングG"),
("事業推進本部 人事G"),
("事業推進本部 メディアG"),
("管理本部"),
("情報システム室"),
("社長室"),
("技術室"),
("監査役"),
("モバイルライフジャパン"),
("インシュアラボ");

INSERT INTO job_models (job) VALUES ("シニアスペシャリスト"),
("スペシャリスト"),
("マネージャー"),
("リーダー"),
("サブリーダー"),
("ジュニア");
