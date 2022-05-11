exports.install = function (Vue, options) {
  Vue.prototype.popWarning = function (msg){
    this.$message({
      showClose: true,
      message: msg,
      type: 'warning'
    })
  }
  Vue.prototype.formatDate = function (row, column, cellValue, index) {
    let date = new Date(parseInt(cellValue)*1000);//时间戳为10位需*1000，如果为13位的话不需乘1000。
    let Y = date.getFullYear() + '-';
    let M = date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) + '-' : date.getMonth() + 1 + '-';
    let D = date.getDate() < 10 ? '0' + date.getDate() + ' ' : date.getDate() + ' ';
    let h = date.getHours() < 10 ? '0' + date.getHours() + ':' : date.getHours() + ':';
    let m = date.getMinutes() < 10 ? '0' + date.getMinutes() + ':' : date.getMinutes() + ':';
    let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
    return Y + M + D + h + m + s;
  }
  Vue.prototype.formatBool = function (row, column, cellValue, index) {
    if (cellValue == true) {
      return "是"
    }
    return "否"
  }
};
