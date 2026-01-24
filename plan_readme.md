# TypeScript 회원 관리 시스템 학습 커리큘럼

## 📋 프로젝트 개요

### 최종 목표

**User Management API Service with Role-Based Access Control**

### 기술 스택

- **Backend**: Node.js + TypeScript + Express
- **Database**: PostgreSQL + Prisma ORM
- **인증**: JWT + Session (하이브리드)
- **보안**: 이메일 인증 + 2FA + OAuth

### 전체 기능

- ✅ 회원가입/로그인 (JWT + Session)
- ✅ 프로필 관리 (CRUD)
- ✅ Role & Permission 시스템 (RBAC)
- ✅ 이메일 인증
- ✅ OAuth 소셜 로그인 (Google, GitHub)
- ✅ 2FA (TOTP)
- ✅ Express Request/Response 타입 확장

### 학습 기간

**총 5주** (35일)

- Week 0: TypeScript 기초 (7일)
- Week 1-4: 프로젝트 개발 (28일)

---

## 🎯 Week 0: TypeScript 기초 집중 학습

> **목표**: 프로젝트 없이 TypeScript 문법만 집중 학습

### ✅ Day 1: 기본 타입

- [x] TypeScript 설치 완료
- [x] 기본 타입 이해 (string, number, boolean)
- [x] 배열과 객체 타입
- [x] any vs unknown
- [x] 함수 타입 정의
- [x] Optional 매개변수와 기본값

### ✅ Day 2: 인터페이스와 타입

- [ ] interface 선언
- [ ] type alias
- [ ] Optional과 Readonly
- [ ] Union과 Intersection
- [ ] interface vs type 차이
- [ ] 타입 확장 (extends)

### ✅ Day 3: 제네릭

- [ ] 제네릭 기본 개념 (<T>)
- [ ] 제네릭 함수
- [ ] 제네릭 인터페이스
- [ ] 제네릭 제약 (extends)
- [ ] keyof 연산자
- [ ] 제네릭 클래스

### ✅ Day 4: 유틸리티 타입

- [ ] Partial, Required, Readonly
- [ ] Pick, Omit
- [ ] Record
- [ ] Exclude, Extract
- [ ] NonNullable
- [ ] ReturnType, Parameters

### ✅ Day 5: 고급 타입

- [ ] 조건부 타입 (Conditional Types)
- [ ] infer 키워드
- [ ] Mapped Types
- [ ] 분산 조건부 타입
- [ ] Template Literal Types
- [ ] 실전 DTO 패턴

### ✅ Day 6: 클래스와 데코레이터

- [ ] 클래스 선언
- [ ] public, private, protected
- [ ] abstract 클래스
- [ ] 인터페이스 구현 (implements)
- [ ] getter/setter
- [ ] 데코레이터 기초 (experimentalDecorators)

### ✅ Day 7: 모듈과 선언 파일

- [ ] export/import 문법
- [ ] Named vs Default export
- [ ] .d.ts 파일 역할
- [ ] declare 키워드
- [ ] Declaration merging
- [ ] @types 패키지 이해
- [ ] Week 0 총 복습

---

## 🚀 Week 1: 프로젝트 환경 설정 + 기본 인증

> **목표**: 프로젝트 구조 설정 및 JWT/Session 인증 구현

### ✅ Day 8: 프로젝트 초기 설정

- [ ] Node.js + Express 프로젝트 생성
- [ ] TypeScript 설정 (tsconfig.json)
- [ ] ESLint + Prettier 설정
- [ ] 폴더 구조 설계
- [ ] 환경 변수 설정 (.env)
- [ ] 개발 스크립트 설정 (nodemon, tsx)

### ✅ Day 9: PostgreSQL + Prisma 설정

- [ ] PostgreSQL 설치 및 연결
- [ ] Prisma 초기화
- [ ] User 모델 스키마 작성
- [ ] Prisma 마이그레이션
- [ ] Prisma Client 생성
- [ ] 기본 CRUD 테스트

### ✅ Day 10: Express 타입 확장 (Part 1)

- [ ] .d.ts 파일 생성 (types/express/index.d.ts)
- [ ] Request 타입 확장 (user, jwt)
- [ ] Response 타입 확장 (success, error)
- [ ] Declaration Merging 적용
- [ ] 타입 안전성 테스트
- [ ] ApiResponse<T> 제네릭 인터페이스

### ✅ Day 11: JWT 인증 구현

- [ ] JWT 라이브러리 설치 (jsonwebtoken)
- [ ] JWT 생성 함수 (sign)
- [ ] JWT 검증 함수 (verify)
- [ ] Access Token 발급
- [ ] Refresh Token 발급
- [ ] authMiddleware 구현

### ✅ Day 12: Session 인증 구현

- [ ] express-session 설정
- [ ] Session 모델 스키마 (Prisma)
- [ ] Session 생성/검증
- [ ] Session 미들웨어
- [ ] JWT + Session 하이브리드
- [ ] 인증 전략 선택 로직

### ✅ Day 13: 회원가입/로그인 API

- [ ] 회원가입 엔드포인트 (/auth/register)
- [ ] 로그인 엔드포인트 (/auth/login)
- [ ] 로그아웃 엔드포인트 (/auth/logout)
- [ ] 비밀번호 해싱 (bcrypt)
- [ ] 입력 검증 (validator)
- [ ] 에러 핸들링

### ✅ Day 14: Week 1 정리 및 테스트

- [ ] API 통합 테스트 (Postman/Thunder Client)
- [ ] 에러 처리 개선
- [ ] 코드 리팩토링
- [ ] README 작성
- [ ] Week 1 회고
- [ ] Week 2 계획 확인

---

## 🔐 Week 2: Role & Permission 시스템

> **목표**: RBAC (Role-Based Access Control) 구현

### ✅ Day 15: Role 모델 설계

- [ ] Role 스키마 작성 (Prisma)
- [ ] Permission 스키마 작성
- [ ] RolePermission 중간 테이블
- [ ] User-Role 관계 설정
- [ ] 마이그레이션 실행
- [ ] Seed 데이터 작성 (admin, user, guest)

### ✅ Day 16: Permission 시스템 구현

- [ ] Permission 타입 정의
- [ ] PermissionService 작성
- [ ] 권한 체크 함수 (hasPermission)
- [ ] 리소스별 권한 정의 (CRUD)
- [ ] 권한 계층 구조
- [ ] 동적 권한 로딩

### ✅ Day 17: Role 미들웨어

- [ ] requireRole 미들웨어
- [ ] 다중 Role 체크
- [ ] Role 기반 라우팅
- [ ] Role 타입 가드
- [ ] 조건부 타입 활용 (RoleRequest<R>)
- [ ] 에러 응답 표준화

### ✅ Day 18: Permission 미들웨어

- [ ] requirePermission 미들웨어
- [ ] 리소스-액션 조합 체크
- [ ] 동적 권한 검증
- [ ] Permission 데코레이터 (선택)
- [ ] 권한 없음 에러 처리
- [ ] 권한 캐싱 전략

### ✅ Day 19: Express 타입 확장 (Part 2)

- [ ] Role 타입 Request 확장
- [ ] Permission 타입 추가
- [ ] 조건부 타입으로 Role별 타입 분기
- [ ] UserWithRole<R> 제네릭 타입
- [ ] PermissionsForRole<R> Mapped Type
- [ ] 타입 안전한 라우트 핸들러

### ✅ Day 20: Admin API 구현

- [ ] 사용자 목록 조회 (admin only)
- [ ] 사용자 Role 변경
- [ ] 사용자 비활성화/삭제
- [ ] 권한 관리 API
- [ ] Role 관리 API
- [ ] 페이지네이션 구현

### ✅ Day 21: Week 2 정리 및 테스트

- [ ] Role & Permission 통합 테스트
- [ ] 권한 시나리오 테스트
- [ ] 에러 케이스 검증
- [ ] 성능 최적화 (N+1 쿼리 해결)
- [ ] 코드 리팩토링
- [ ] Week 2 회고

---

## 📧 Week 3: 이메일 인증 + 2FA

> **목표**: 보안 강화 (이메일 인증, 2단계 인증)

### ✅ Day 22: 이메일 서비스 설정

- [ ] Nodemailer 설정
- [ ] SMTP 설정 (Gmail/SendGrid)
- [ ] 이메일 템플릿 작성
- [ ] 이메일 발송 함수
- [ ] 비동기 큐 (선택, Bull)
- [ ] 이메일 로깅

### ✅ Day 23: 이메일 인증 구현

- [ ] 인증 토큰 생성 (UUID/JWT)
- [ ] 인증 링크 생성
- [ ] 인증 이메일 발송
- [ ] 이메일 인증 엔드포인트
- [ ] 인증 상태 확인 미들웨어
- [ ] 재발송 기능

### ✅ Day 24: 비밀번호 재설정

- [ ] 비밀번호 재설정 요청 API
- [ ] 재설정 토큰 생성
- [ ] 재설정 이메일 발송
- [ ] 재설정 확인 API
- [ ] 새 비밀번호 설정
- [ ] 토큰 만료 처리

### ✅ Day 25: 2FA 설정 (TOTP)

- [ ] speakeasy 라이브러리 설치
- [ ] 2FA Secret 생성
- [ ] QR 코드 생성 (qrcode)
- [ ] 2FA 활성화 API
- [ ] TOTP 토큰 검증
- [ ] 백업 코드 생성

### ✅ Day 26: 2FA 인증 플로우

- [ ] 로그인 시 2FA 체크
- [ ] 2FA 코드 입력 엔드포인트
- [ ] 2FA 필수 미들웨어
- [ ] 신뢰된 디바이스 (선택)
- [ ] 2FA 비활성화 API
- [ ] 복구 코드 사용

### ✅ Day 27: Express 타입 확장 (Part 3)

- [ ] 2FA 상태 Request 타입
- [ ] 이메일 인증 상태 타입
- [ ] 조건부 응답 타입
- [ ] 보안 컨텍스트 타입
- [ ] 인증 레벨 타입 정의
- [ ] 타입 가드 함수

### ✅ Day 28: Week 3 정리 및 테스트

- [ ] 이메일 인증 플로우 테스트
- [ ] 2FA 전체 시나리오 테스트
- [ ] 보안 취약점 점검
- [ ] 에러 처리 개선
- [ ] 코드 리팩토링
- [ ] Week 3 회고

---

## 🔗 Week 4: OAuth + 프로필 관리

> **목표**: 소셜 로그인 및 사용자 프로필 기능

### ✅ Day 29: OAuth 기초 설정

- [ ] Passport.js 설치
- [ ] OAuth 전략 개념 이해
- [ ] Google OAuth 앱 생성
- [ ] GitHub OAuth 앱 생성
- [ ] 환경 변수 설정
- [ ] Callback URL 설정

### ✅ Day 30: Google OAuth 구현

- [ ] passport-google-oauth20 설정
- [ ] Google 전략 구현
- [ ] OAuth 콜백 핸들러
- [ ] 기존 계정 연결
- [ ] 신규 사용자 생성
- [ ] OAuth 사용자 타입 정의

### ✅ Day 31: GitHub OAuth 구현

- [ ] passport-github2 설정
- [ ] GitHub 전략 구현
- [ ] OAuth 프로바이더 추상화
- [ ] 다중 OAuth 프로바이더 지원
- [ ] OAuth 연결 해제
- [ ] OAuth 계정 관리 API

### ✅ Day 32: 프로필 관리 API

- [ ] 프로필 조회 API
- [ ] 프로필 수정 API
- [ ] 프로필 이미지 업로드 (multer)
- [ ] 비밀번호 변경 API
- [ ] 계정 삭제 API
- [ ] 활동 로그 조회

### ✅ Day 33: 파일 업로드 처리

- [ ] Multer 설정
- [ ] 이미지 검증 (타입, 크기)
- [ ] 파일 저장 (로컬/S3)
- [ ] 썸네일 생성 (sharp, 선택)
- [ ] 파일 타입 정의
- [ ] Express.Multer.File 타입 확장

### ✅ Day 34: 페이지네이션 & 필터링

- [ ] 페이지네이션 타입 정의
- [ ] 페이지네이션 미들웨어
- [ ] 필터링 쿼리 파싱
- [ ] 정렬 기능
- [ ] 검색 기능
- [ ] Response에 메타 정보 추가

### ✅ Day 35: Week 4 정리 및 테스트

- [ ] OAuth 플로우 통합 테스트
- [ ] 프로필 CRUD 테스트
- [ ] 파일 업로드 테스트
- [ ] 전체 API 문서화 (Swagger 선택)
- [ ] 코드 리팩토링
- [ ] Week 4 회고

---

## 🎓 Week 5: 최종 완성 및 배포

> **목표**: 프로젝트 마무리, 테스트, 문서화, 배포

### ✅ Day 36: 통합 테스트

- [ ] Jest 설정
- [ ] 단위 테스트 작성
- [ ] 통합 테스트 작성
- [ ] E2E 테스트 (선택)
- [ ] 테스트 커버리지 확인
- [ ] CI 설정 (GitHub Actions, 선택)

### ✅ Day 37: 에러 처리 표준화

- [ ] 커스텀 에러 클래스
- [ ] 전역 에러 핸들러
- [ ] 에러 코드 enum
- [ ] 에러 로깅 (Winston)
- [ ] 에러 응답 타입 통일
- [ ] 프로덕션 에러 처리

### ✅ Day 38: 로깅 & 모니터링

- [ ] Winston 로거 설정
- [ ] 로그 레벨 관리
- [ ] 요청 로깅 미들웨어
- [ ] 에러 로깅
- [ ] 성능 모니터링 (선택)
- [ ] 로그 로테이션

### ✅ Day 39: 보안 강화

- [ ] Helmet.js 적용
- [ ] CORS 설정
- [ ] Rate Limiting (express-rate-limit)
- [ ] SQL Injection 방어
- [ ] XSS 방어
- [ ] CSRF 토큰 (선택)

### ✅ Day 40: API 문서화

- [ ] Swagger/OpenAPI 설정
- [ ] API 엔드포인트 문서화
- [ ] 타입 스키마 문서화
- [ ] 예제 요청/응답
- [ ] Postman Collection
- [ ] README 업데이트

### ✅ Day 41: 배포 준비

- [ ] 프로덕션 환경 설정
- [ ] Docker 설정 (선택)
- [ ] 데이터베이스 마이그레이션 전략
- [ ] 환경 변수 관리
- [ ] 빌드 스크립트
- [ ] Health Check 엔드포인트

### ✅ Day 42: 최종 정리 및 회고

- [ ] 전체 기능 최종 테스트
- [ ] 성능 최적화
- [ ] 코드 리뷰 및 리팩토링
- [ ] 프로젝트 문서 완성
- [ ] 학습 회고록 작성
- [ ] CTF Toolkit 계획 수립

---

## 📊 학습 체크리스트 요약

### TypeScript 핵심 개념

- [ ] 기본 타입 시스템
- [ ] 인터페이스와 타입
- [ ] 제네릭
- [ ] 유틸리티 타입
- [ ] 고급 타입 (조건부, Mapped)
- [ ] 클래스와 데코레이터
- [ ] 모듈과 선언 파일

### Express + TypeScript

- [ ] Express Request/Response 타입 확장
- [ ] Declaration Merging
- [ ] 미들웨어 타입 정의
- [ ] 라우트 핸들러 타입
- [ ] 에러 핸들러 타입
- [ ] 제네릭 응답 패턴

### Prisma + TypeScript

- [ ] Generated Types 활용
- [ ] Relation 타입
- [ ] DTO 타입 변환
- [ ] 쿼리 결과 타입
- [ ] Partial 업데이트
- [ ] Transaction 타입

### 인증 & 보안

- [ ] JWT 타입 안전성
- [ ] Session 타입 정의
- [ ] Role 타입 시스템
- [ ] Permission 타입 검증
- [ ] 2FA 타입 정의
- [ ] OAuth 사용자 타입

### 실전 패턴

- [ ] Service Layer 패턴
- [ ] Repository 패턴
- [ ] DTO 패턴
- [ ] 미들웨어 체이닝
- [ ] 에러 핸들링 패턴
- [ ] 타입 가드 활용

---

## 🎯 최종 결과물

### 완성된 API 엔드포인트

**인증 (Auth)**

- POST /auth/register - 회원가입
- POST /auth/login - 로그인
- POST /auth/logout - 로그아웃
- POST /auth/refresh - 토큰 갱신
- POST /auth/verify-email - 이메일 인증
- POST /auth/forgot-password - 비밀번호 재설정 요청
- POST /auth/reset-password - 비밀번호 재설정

**2FA**

- POST /auth/2fa/setup - 2FA 설정
- POST /auth/2fa/verify - 2FA 인증
- POST /auth/2fa/disable - 2FA 비활성화

**OAuth**

- GET /auth/google - Google 로그인
- GET /auth/google/callback - Google 콜백
- GET /auth/github - GitHub 로그인
- GET /auth/github/callback - GitHub 콜백

**프로필 (Profile)**

- GET /profile - 프로필 조회
- PUT /profile - 프로필 수정
- POST /profile/avatar - 프로필 이미지 업로드
- PUT /profile/password - 비밀번호 변경
- DELETE /profile - 계정 삭제

**관리자 (Admin)**

- GET /admin/users - 사용자 목록
- PUT /admin/users/:id/role - Role 변경
- DELETE /admin/users/:id - 사용자 삭제
- GET /admin/roles - Role 목록
- POST /admin/permissions - Permission 생성

### 주요 타입 정의 파일

```
src/types/
├── express/
│   └── index.d.ts          # Express 타입 확장
├── models/
│   ├── user.ts
│   ├── role.ts
│   ├── permission.ts
│   ├── session.ts
│   └── oauth.ts
├── api/
│   ├── request.ts
│   ├── response.ts
│   └── pagination.ts
└── utils/
    ├── jwt.ts
    └── validator.ts
```

### 학습 성과

- ✅ TypeScript 타입 시스템 완전 이해
- ✅ Express와 TypeScript 통합 마스터
- ✅ Prisma ORM 타입 활용
- ✅ RBAC 시스템 설계 및 구현
- ✅ 실무 인증/인가 패턴
- ✅ 타입 안전한 API 개발
- ✅ 보안 Best Practices

---

## 📝 다음 단계: CTF Toolkit

이 프로젝트 완성 후:

1. TypeScript 기초 완벽 습득
2. Express 타입 확장 경험
3. 복잡한 타입 시스템 설계 능력
4. 실무 패턴 이해

→ **CTF Toolkit 프로젝트**에 응용 가능!

---

**작성일**: 2025-12-16
**예상 완료일**: 2026-01-20 (5주)
**최종 목표**: TypeScript 마스터 + 실전 회원 관리 시스템 완성
