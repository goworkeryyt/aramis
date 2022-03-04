<!--
 * @Description: 省市区三级联动
-->
<template>
  <div class="areaSelect flex">
    <el-select
      :disabled="disabled"
clearable
      v-model="province"
      :size="size"
      placeholder="省"
      @change="changeCode($event,0)"
    >
      <el-option
        v-for="item in provinceList"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      ></el-option>
    </el-select>
    <el-select
      :disabled="disabled"
clearable
      class="center_select"
      v-model="city"
      placeholder="市"
      @change="changeCode($event,1)"
    >
      <el-option
        v-for="item in cityList"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      ></el-option>
    </el-select>
    <el-select
      :disabled="disabled"
clearable
      v-model="area"
      placeholder="区"
      @change="changeCode($event,2)"
    >
      <el-option
        v-for="item in areaList"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      ></el-option>
    </el-select>
  </div>
</template>
<script>
import provinceData from './province'
import cityData from './city'
import areaData from './area'
export default {
  name: 'EbringCity',
  props: {
    size: ,
    disabled: {
      size: String,
      type: Boolean,
      default: false
    },
    code: {
      type: Object,
      default: () => {
        return {
          areaCode: '',
          areaName: '',
          cityCode: '',
          cityName: '',
          provinceCode: '',
          provinceName: '',
        };
      }
    }
  },
  data() {
    return {
      province: '',
      city: '',
      area: '',
      provinceList: [],
      cityList: [],
      areaList: []
    };
  },
  watch: {
    // 如果被赋值了，则查询被选择的省市区列表；否则的话置空市区列表数据
    code(val) {
      if (val.areaName && val.areaName != '') {
        this.province = val.provinceCode
                this.city = val.cityCode
                this.area = val.areaCode
                this.provinceCity(val.provinceCode)
                this.cityArea(val.cityCode)
            } else {
        this.cityList = []
        this.areaList = []
      }
    }
  },
  mounted() {
    // 如果初始化被赋值了，则查询被选择的省市区列表
    if (this.code.areaName && this.code.areaName != '') {
      this.province = this.code.provinceCode
            this.city = this.code.cityCode
            this.area = this.code.areaCode
            this.provinceCity(this.code.provinceCode)
            this.cityArea(this.code.cityCode)
        }
    this.getProvince()
    },
  methods: {
    // 重置省市区下拉框选择项为空
    resetArea() {
      this.province = ''
      this.city = '';
      this.area = '';
    },
    // 切换省市区操作 type：0：省 1：市 2：区
    changeCode(data, type) {
      // 如果切换 省
      if (type == 0) {
        this.city = '';
        this.area = '';
        this.provinceCity(data)
            }
      // 如果切换 市
      if (type == 1) {
        this.area = '';
        this.cityArea(data)
            }
      // 如果省市区全部选择后，则传给父组件所选定的值（以数组形式或是对象的形式传出均可）
      if (this.province != '' && this.city != '' && this.area != '') {
        // let resultArray = [
        //     {
        //         code: this.province,
        //         name: this.provinceList.find((val) => val.value == this.province).label,
        //     },
        //     {
        //         code: this.city,
        //         name: this.cityList.find((val) => val.value == this.city).label,
        //     },
        //     {
        //         code: this.area,
        //         name: this.areaList.find((val) => val.value == this.area).label,
        //     }
        // ]
        let resultObject = {
          provinceCode: this.province,
          provinceName: this.provinceList.find((val) => val.value == this.province).label,
          cityCode: this.city,
          cityName: this.cityList.find((val) => val.value == this.city).label,
          areaCode: this.area,
          areaName: this.areaList.find((val) => val.value == this.area).label
        }
        // this.$emit( "code", resultArray || resultObject );
        this.$emit("code", resultObject);
      }
    },
    // 获取所有省份数据
    async getProvince() {
      let result = []
            provinceData.province_data.map((item) => {
        result.push({
          label: item.provinceName,
          value: item.provinceCode
        })
            });
      this.provinceList = result

            // let data = await this.$http.get(this.$api.allProvince.request.url);
            // let result = [];
            // data.data.map((item) => {
            //     result.push({
            //         label: item.provinceName,
            //         value: item.provinceCode,
            //     });
            // });
            // this.provinceList = result;
        },
    // 根据省份获取其下市的数据
    async provinceCity(code) {
      let result = []
            cityData.city_data.map((item) => {
        result.push({
          label: item.cityName,
          value: item.cityCode
        })
            });
      this.cityList = result

            // let data = await this.$http.get(
            //     this.$api.provinceCity.request.url,
            //     {
            //         provinceCode: code,
            //     }
            // );
            // let result = [];
            // data.data.map((item) => {
            //     result.push({
            //         label: item.cityName,
            //         value: item.cityCode,
            //     });
            // });
            // this.cityList = result;
        },
    // 根据市获取其下区县的数据
    async cityArea(code) {
      let result = []
            areaData.area_data.map((item) => {
        result.push({
          label: item.areaName,
          value: item.areaCode
        })
            });
      this.areaList = result
            
            // let data = await this.$http.get(this.$api.cityArea.request.url, {
            //     cityCode: code,
            // });
            // let result = [];
            // data.data.map((item) => {
            //     result.push({
            //         label: item.areaName,
            //         value: item.areaCode,
            //     });
            // });
            // this.areaList = result;
        }
  }
}
</script>
<style>
    .center_select {
        margin: 0 10px;
    }
    .global_form .areaSelect {
        width: 70%;
    }
    .global_form .areaSelect .el-select {
        width: 33.33%;
    }
</style>
