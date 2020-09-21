<template>
  <div class="content">
    <div class="md-layout">
      <div class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33">
        <benchmark-chart :type="'producer'" :num="100" :size="1000" background-color="green" />
      </div>
      <div class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33">
        <benchmark-chart :type="'producer'" :num="1000" :size="1000" background-color="red" />
      </div>
      <div class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33">
        <benchmark-chart :type="'producer'" :num="10000" :size="1000" background-color="orange" />
      </div>
    </div>
  </div>
</template>

<script>
import BenchmarkChart from "../components/BenchmarkChart.vue";

export default {
  components: {
    BenchmarkChart
  }
};
</script>
