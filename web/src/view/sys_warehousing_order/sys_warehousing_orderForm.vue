<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="车牌号:">
                <el-input v-model="formData.numbersOfCar" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="货物类型:">
                  <el-input-number v-model="formData.typeOfGoods" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="公司名字:">
                <el-input v-model="formData.companyName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="司机名字:">
                <el-input v-model="formData.driverName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="司机电话号:">
                <el-input v-model="formData.driver_phone" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createWarehousingOrder,
    updateWarehousingOrder,
    findWarehousingOrder
} from "@/api/sys_warehousing_order";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "WarehousingOrder",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            numbersOfCar:"",
            typeOfGoods:0,
            companyName:"",
            driverName:"",
            driver_phone:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createWarehousingOrder(this.formData);
          break;
        case "update":
          res = await updateWarehousingOrder(this.formData);
          break;
        default:
          res = await createWarehousingOrder(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findWarehousingOrder({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rewo
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>