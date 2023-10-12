<template>
  <div class="d-flex">
    <select id="form.selectCity" class="form-select me-2" style="width: 30%" place v-model="selectedCity">
      <option value="" disabled selected>Chọn tỉnh thành</option>
      <option v-for="city in cities" :key="city.Id" :value="city.Id">{{ city.Name }}</option>
    </select>

    <select class="form-select me-2" style="width: 30%" v-model="selectedDistrict">
      <option value="" disabled selected>Chọn quận huyện</option>
      <option v-for="district in districts" :key="district.Id" :value="district.Id">{{ district.Name }}</option>
    </select>

    <select class="form-select" style="width: 30%" v-model="selectedWard">
      <option value="" disabled selected>Chọn phường xã</option>
      <option v-for="ward in wards" :key="ward.Id" :value="ward.Id">{{ ward.Name }}</option>
    </select>
  </div>

</template>
<script>
import axios from 'axios'

export default {
  name: 'PlaceSelection',
  props: ['value'],
  data () {
    return {
      cities: [],
      districts: [],
      wards: [],
      selectedCity: '',
      selectedDistrict: '',
      selectedWard: ''
    }
  },
  async created () {
    try {
      const response = await axios.get('https://raw.githubusercontent.com/kenzouno1/DiaGioiHanhChinhVN/master/data.json')
      this.cities = response.data
    } catch (error) {
      console.error(error)
    }
  },
  watch: {
    selectedCity (newVal) {
      const city = this.cities.find((c) => c.Id === newVal)
      if (city) {
        this.districts = city.Districts
        this.selectedDistrict = ''
        this.wards = []
      }
      this.emitChange()
    },
    selectedDistrict (newVal) {
      const city = this.cities.find((c) => c.Id === this.selectedCity)
      if (city) {
        const district = city.Districts.find((d) => d.Id === newVal)
        if (district) {
          this.wards = district.Wards
          this.selectedWard = ''
        }
      }
      this.emitChange()
    },
    selectedWard (newVal) {
      this.emitChange()
    }
  },
  methods: {
    emitChange () {
      this.$emit('place-selected', {
        city: this.cities.find((c) => c.Id === this.selectedCity),
        district: this.districts.find((d) => d.Id === this.selectedDistrict),
        ward: this.wards.find((w) => w.Id === this.selectedWard)
      })
    }
  }
}
</script>
