# 7. Модель данных

Здесь описываются сущности, которые будут храниться в PostgreSQL, Mongo

# Postgres

## Users (Пользователи)
- id
- username
- email
- password (hash)

## User_info (Дополнительная информация о пользователях)
- id
- user_id
- city
- about
- role
- class
- subject_ids

## User_settings (Настройки пользователя)
- id
- user_id
- is_2fa_enabled
- is_notifications_enabled

## Groups (Учебные группы)
- id
- title
- description
- subject_id
- students_count
- tg_group_link
- tg_chat_id
- tutor_id

## User_groups (Связь учеников и групп)
- id
- user_id
- group_id

## Schedules (Расписания)
- id
- group_id
- repeat_type (once, weekly, monthly)
- start_date
- end_date
- is_active

## Schedule_slots (Слоты, вкюченные в расписание)
- id
- schedule_id
- weekday
- start_time
- end_time
- duration

## Boards (Доски)
- id
- snapshot_version
- group_id

## Subjects (Предметы)
- id
- code
- title
- type (ЕГЭ, ОГЭ, повышение успеваемости)

## Webinars (Вебинары)
- id
- subject_id
- group_id
- schedule_slot_id
- status (planned, active, ended)
- meet_link
- board_id

## Webinar_participants (Ученики с доступом к вебинару)
- id
- webinar_id
- user_id

# Mongo

## Homeworks (Домашние работы)
- _id
- user_ids
- group_id
- type (test, homework, video)
- start_time
- end_time
- status (opened, completed, closed)
- is_deadline_complied
- exercises_count
- correct_answers_count

## Homework_content (Содержание домашней работы)
- _id
- homework_id
- video_link
- exercises: [{
    - _id
    - exercise_id
    - condition
    - image_link
    - answer_type (single, multiple, text, image, input)
    - variable_answers
    - answer
}]

## Board_snapshots (Скриншоты доски)
- _id
- board_id
- version
- objects: [{
    - id
    - object_id
    - type (line, geometric_figure, text, image)
    - x
    - y
    - width
    - height
    - color
    - stroke_width
    - text
    - image_link
    }
]
- created_at

## Board_events (События на доске)
- _id
- board_id
- type (add, delete)
- object: {
    - _id
    - object_id
    - type
    - x
    - y
    - width
    - height
    - color
    - stroke_width
    - text
    - media_link
  }