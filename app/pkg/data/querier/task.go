package querier

import (
    "colab/app/pkg/data/example"
    "gorm.io/gen"
)

type TaskQuerier interface {

    // QueryTaskByAttr
    // SELECT t.*
    // FROM colab_task t
    // {{ if len(attrs) > 0 }}
    //
    //	JOIN (SELECT cea.biz_id task_id
    //	        {{ for _, attr1 := range attrs }}
    //	           , max(if(cea.code = @attr1.Code, cea.value, NULL))   AS @@attr1.Code
    //	        {{ end }}
    //	      FROM colab_entity_attribute cea
    //	      WHERE cea.biz = 'task'
    //	      GROUP BY task_id) attrs ON t.id = attrs.task_id
    //
    // {{ end }}
    // WHERE t.deleted_at IS NULL
    // {{ if len(attrs) > 0 }}
    //
    //	{{ for _, attr2 := range attrs }}
    //	   AND attrs.@@attr2.Code = @attr2.Value
    //	{{ end }}
    //
    // {{ end }}
    QueryTaskByAttr(attrs []example.Attr) ([]gen.T, error)
}
