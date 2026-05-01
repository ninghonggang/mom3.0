# Report报表模块设计文档

## 1. 模块概述

Report报表模块是SFMS3.0的统一报表管理模块，负责数据可视化大屏配置、报表设计、SQL数据查询等核心功能，为企业提供数据分析和决策支持能力。

**路径**: `win-module-report`

**子模块结构**:
- `win-module-report-api` - API接口定义
- `win-module-report-biz` - 业务实现

---

## 2. 模块职责

| 子模块 | 职责 |
|--------|------|
| GoView大屏 | 基于Vue3+GoView的数据大屏配置与展示 |
| AJ-Report报表 | 集成积木报表（Jeecg-boot子项目）的中国式复杂报表 |
| UReport报表 | 集成UReport2表达式报表引擎 |
| 报表Token | 统一认证与权限集成 |

### 2.1 GoView大屏模块

GoView是前端大屏可视化组件，本模块提供后端API支持：

- **项目管理**: 大屏项目的CRUD操作
- **数据源配置**: 支持SQL查询和HTTP接口两种数据获取方式
- **预览与发布**: 项目状态管理（已发布/未发布）

### 2.2 AJ-Report积木报表

集成Jeecg-boot的积木报表引擎，提供中国式复杂报表能力：

- **报表设计器**: 拖拽式报表设计
- **数据源集成**: 统一认证和数据源管理
- **表达式扩展**: 自定义表达式函数（如NumberFormat）

### 2.3 报表Token认证

通过`JmReportTokenServiceImpl`实现报表系统与主系统的SSO集成：

- 复用OAuth2TokenApi进行Token验证
- 集成SecurityProperties安全配置

---

## 3. 核心类/接口

### 3.1 GoView项目控制器

**Controller**: `GoViewProjectController`
- 路径: `controller/goview/GoViewProjectController.java`
- 职责: 大屏项目的CRUD、导入导出

```java
@Tag(name = "管理后台 - GoView 项目")
@RestController
@RequestMapping("/report/go-view/project")
public class GoViewProjectController {
    // POST   /report/go-view/project/create    - 创建项目
    // PUT    /report/go-view/project/update    - 更新项目
    // DELETE /report/go-view/project/delete   - 删除项目
    // GET    /report/go-view/project/get      - 获取项目
    // GET    /report/go-view/project/my-page  - 分页查询
}
```

### 3.2 GoView数据服务

**Controller**: `GoViewDataController`
- 路径: `controller/goview/GoViewDataController.java`
- 职责: 数据查询接口

```java
@Tag(name = "管理后台 - GoView 数据")
@RestController
@RequestMapping("/report/go-view/data")
public class GoViewDataController {
    // POST /report/go-view/data/get-by-sql   - SQL查询数据
    // POST /report/go-view/data/get-by-http - HTTP查询数据
}
```

**Service接口**: `GoViewDataService`
```java
public interface GoViewDataService {
    GoViewDataRespVO getDataBySQL(String sql);
}
```

### 3.3 GoView数据模型

**DO**: `GoViewProjectDO`
- 路径: `dal/dataobject/goview/GoViewProjectDO.java`
- 表名: `report_go_view_project`

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 编号 |
| name | String | 项目名称 |
| picUrl | String | 预览图片URL |
| content | String | JSON配置内容 |
| status | Integer | 发布状态(0-已发布/1-未发布) |
| remark | String | 备注 |

**VO**: `GoViewDataRespVO`
```java
public class GoViewDataRespVO {
    private List<String> dimensions;  // 数据维度
    private List<Map<String, Object>> source;  // 明细数据
}
```

### 3.4 积木报表配置

**Configuration**: `JmReportConfiguration`
- 路径: `framework/jmreport/config/JmReportConfiguration.java`
- 职责: 扫描积木报表包，配置Token服务

**Express扩展**: `JmExpressCustomImpl`
- 实现`IJmExpressCustom`接口
- 注册自定义表达式函数（如NumberFormat）

---

## 4. 数据结构

### 4.1 GoView项目表

```
表名: report_go_view_project
主键: id (自增)

字段:
- id          BIGINT       编号
- name        VARCHAR(100) 项目名称
- pic_url     VARCHAR(500) 预览图片
- content     TEXT         JSON配置
- status      INT          状态(0-已发布/1-未发布)
- remark      VARCHAR(500) 备注
- create_time DATETIME     创建时间
- update_time DATETIME     更新时间
```

### 4.2 VO对象

| VO类 | 用途 |
|------|------|
| GoViewProjectCreateReqVO | 创建项目请求 |
| GoViewProjectUpdateReqVO | 更新项目请求 |
| GoViewProjectRespVO | 项目响应 |
| GoViewDataGetBySqlReqVO | SQL查询请求 |
| GoViewDataRespVO | 数据响应 |

---

## 5. 数据流向

```
[前端大屏] 
    |
    |-- SQL查询 --> GoViewDataController.getDataBySQL()
    |                        |
    |                        v
    |              GoViewDataServiceImpl.getDataBySQL()
    |                        |
    |                        v
    |              直接执行SQL返回JSON数据
    |
    |-- HTTP接口 --> GoViewDataController.getByHttp()
    |                        |
    |                        v
    |              业务逻辑返回模拟/外部数据
    |
    |-- 项目管理 --> GoViewProjectController
                           |
                           v
                 GoViewProjectService --> MySQL
```

---

## 6. 关键技术实现

### 6.1 技术栈

- **前端组件**: GoView（Vue3大屏可视化库）
- **报表引擎**: AJ-Report（积木报表）、UReport2
- **数据存储**: MySQL（项目配置）、直连业务库（SQL查询）
- **认证集成**: OAuth2 Token + Security Framework

### 6.2 SQL数据查询

直接执行用户传入的SQL语句，返回JSON格式结果：

```java
public GoViewDataRespVO getDataBySQL(String sql) {
    // SQL查询结果转换为dimensions + source结构
    // dimensions: 列名列表
    // source: Map列表，每行数据
}
```

### 6.3 权限控制

使用`@PreAuthorize("@ss.hasPermission('report:go-view-project:create')")`注解实现细粒度权限控制。

---

## 7. 集成关系

```
Report模块
    |
    +-- 依赖 Infra模块（OAuth2TokenApi, SecurityProperties）
    |
    +-- 依赖 System模块（用户、权限数据）
    |
    +-- 被各业务模块调用（生产、仓库、采购等的数据展示）
```

---

## 8. 错误码

| 错误码 | 说明 |
|--------|------|
| 1-003-000-000 | GoView项目不存在 |
