import {defineStore} from "pinia";
//report api
import {useReportApi} from "/@/api/report";
import {ElMessage} from "element-plus";

const reportApi = useReportApi()
export const useReportStore = defineStore("reportStore", {
    state: () => ({

        //当前数据库的类型和数据库列表
        dbInfo: {
            db_type: '',
            database_list: [],
        },
        mysqlTable: [] as MysqlTable[],
        mysqlColumn: [] as MysqlColumn[],
        mysqlColumnTypeMap: new Map(),
        mysqlColumnChineseNameMap: new Map(),

        sqliteTable: [] as SqliteTable[],
        sqliteColumn: [] as SqliteColumn[],
        sqliteColumnTypeMap: new Map(),

    }),
    actions: {
        async getDB(params?: object) {
            const res = await reportApi.getDBApi()
            if (res.code === 0) {
                this.dbInfo = res.data
            }
        },
        //加载库表，参数：{"database":"xxxx}
        async getTables(params?: object) {
            const res = await reportApi.getTablesApi(params)
            if (res.code === 0) {
                // ElMessage.success(res.msg)
                switch (this.dbInfo.db_type) {
                    case "mysql":
                        this.mysqlTable = res.data
                        break
                    case "sqlite":
                        this.sqliteTable = res.data
                        break
                }
            }
        },
        //加载字段，参数：{"database":"xxxx","table_name":"xxx}
        async getColumn(params?: object) {
            const res = await reportApi.getColumnApi(params)
            if (res.code === 0) {
                // ElMessage.success(res.msg)
                switch (this.dbInfo.db_type) {
                    case "mysql":
                        this.mysqlColumn = res.data
                        this.mysqlColumnTypeMapHandler()
                        break
                    case "sqlite":
                        this.sqliteColumn = res.data
                        this.sqliteColumnTypeMapHandler()
                        break
                }
            }
        },
        //处理字段类型
        sqliteColumnTypeMapHandler() {
            const colMap = new Map()
            this.sqliteColumn.forEach((value, index, array) => {
                colMap.set(value.name, value.type)
            })
            this.sqliteColumnTypeMap = colMap
        },
        mysqlColumnTypeMapHandler() {
            const colMap = new Map()
            const colChineseNameMap = new Map()
            this.mysqlColumn.forEach((value, index, array) => {
                colMap.set(value.column_name, value.data_type)
                colChineseNameMap.set(value.column_name, value.column_comment)
            })
            this.mysqlColumnTypeMap = colMap
            this.mysqlColumnChineseNameMap = colChineseNameMap
        }
    }

})