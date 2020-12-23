# study_app

## 仕様
- 日々の学習結果を記録できるアプリ
    - 学習対象の本などを作成/参照/編集/削除
        - 本一覧ページ
        - 本作成ページ
    - 学習データを作成/編集/編集/削除
        - 学習データ一覧ページ
    - ユーザーを作成/参照/編集/削除
        - ユーザー一覧ページ
        - ユーザー作成ページ

## テーブル構造
- User
    - id
    - name
    - email
    - image_url
    - created_at
    - updated_at
- Book
    - id
    - genre_id
    - title
    - author
    - publisher
    - image_url
    - urgency_number
    - priority_number
    - published_at
    - created_at
    - updated_at
- Genre
    - id
    - name
- Study
    - id
    - user_id
    - result
    - executed_at
    - created_at
    - updated_at
- StudyPlanItem
    - id
    - study_id
    - book_id
    - memo
- StudyResultItem
    - id
    - study_id
    - book_id
    - memo
- UserBookItem
    - id
    - user_id
    - book_id
    - status
    - progress_rate

## 技術スタック
### フロントエンド
- HTML/CSS
- JavaScript/TypeScript
- React/Next.js
- Bootstrap or Material-UI
### バックエンド
- Golang
- echo
- gorm
- OpenAPI
### DB
- PostgreSQL or MySQL
### インフラ
- AWS
    - ECS/Fargate or EKS/Fargate
    - ECR
    - RDS
    - S3
    - VPC
    - ALB
    - Route53
### その他
- Docker / docker-compose
- Kubernetes
- github actions
