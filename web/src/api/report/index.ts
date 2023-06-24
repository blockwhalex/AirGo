import service from '/@/utils/request';

export function useReportApi() {
    return {
        //getDB
        getDBApi: () => {
            return service({
                url: "/report/getDB",
                method: "get",
            })
        },
        // getTables
        getTablesApi: (data?: object) => {
            return service({
                url: "/report/getTables",
                method: "post",
                data
            })
        },
        // getColumn
        getColumnApi: (data?: object) => {
            return service({
                url: "/report/getColumn",
                method: "post",
                data
            })
        },
        //提交
        submitReportApi: (data?: object) => {
            return service({
                url: "/report/reportSubmit",
                method: "post",
                data
            })
        },
    }
}