# gorm
> gorm中文文档： https://learnku.com/docs/gorm/v2

值得一提的是gorm的v1与v2在很多地方都不兼容：
```
// v1
import ( 
    "github.com/jinzhu/gorm"    
    // dialects
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

// v2
import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)
```

