iot-platform/
│main.go                  # 主程序入口文件
│
├── config/                       # 配置文件目录
│   ├── config.ini               # 主配置文件
│   ├── config.go                # 配置结构体定义
│   └── mqtt/                    # MQTT 配置
│       └── mosquitto.conf       # MQTT 配置文件
│
├── global/                       # 全局变量
│   └── app.go                   # 全局变量定义
│
├── internal/                     # 内部应用代码
│   ├── auth/                    # 认证与授权
│   │   ├── jwt.go              # JWT 实现
│   │   ├── middleware.go       # 认证中间件
│   │   └── permission.go       # 权限控制
│   │
│   ├── controller/             # 控制器层
│   │   ├── device.go          # 设备控制器
│   │   ├── user.go            # 用户控制器
│   │   └── gateway.go         # 网关控制器
│   │
│   ├── model/                 # 数据模型层
│   │   ├── device.go         # 设备模型
│   │   ├── user.go           # 用户模型
│   │   └── gateway.go        # 网关模型
│   │
│   ├── service/              # 业务逻辑层
│   │   ├── device.go        # 设备服务
│   │   ├── user.go          # 用户服务
│   │   └── gateway.go       # 网关服务
│   │
│   └── repository/          # 数据访问层
│       ├── mysql/          # MySQL 实现
│       └── redis/          # Redis 实现
│
├── pkg/                    # 可复用的包
│   ├── cache/             # 缓存实现
│   │   ├── redis.go      # Redis 缓存
│   │   └── memory.go     # 内存缓存
│   │
│   ├── database/         # 数据库操作
│   │   ├── mysql.go     # MySQL 连接
│   │   └── migration.go # 数据库迁移
│   │
│   ├── logger/          # 日志包
│   │   ├── logger.go    # 日志实现
│   │   └── options.go   # 日志配置
│   │
│   ├── protocol/        # 协议实现
│   │   ├── mqtt/       # MQTT 协议
│   │   │   ├── client.go
│   │   │   └── handler.go
│   │   ├── tcp/        # TCP 协议
│   │   │   ├── server.go
│   │   │   └── handler.go
│   │   └── modbus/     # Modbus 协议
│   │       ├── client.go
│   │       └── handler.go
│   │
│   └── utils/          # 工具函数
│       ├── encrypt.go  # 加密工具
│       └── validator.go # 验证工具
│
├── public/             # 静态文件
│   ├── images/        # 图片文件
│   ├── js/           # JavaScript 文件
│   └── css/          # CSS 文件
│
├── runtime/           # 运行时数据
│   ├── logs/         # 日志文件
│   │   ├── app.log
│   │   └── error.log
│   └── upload/       # 上传文件
│
├── scripts/          # 脚本文件
│   ├── deploy.sh    # 部署脚本
│   └── backup.sh    # 备份脚本
│
├── test/            # 测试文件
│   ├── api/        # API 测试
│   └── unit/       # 单元测试
│
├── .gitignore      # Git 忽略文件
├── go.mod          # Go 模块文件
├── go.sum          # Go 依赖校验
└── README.md       # 项目说明文件