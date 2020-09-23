<template>
    <div class="content">
        <div class="md-layout">
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart :type="'consumer'" :num="10000" :size="1000" />
            </div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart
                    :type="'consumer'"
                    :num="100000"
                    :size="1000"
                />
            </div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart
                    :type="'consumer'"
                    :num="1000000"
                    :size="1000"
                />
            </div>
        </div>
        <div class="md-layout">
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart
                    :type="'consumer'"
                    :num="10000000"
                    :size="1000"
                />
            </div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            ></div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            ></div>
        </div>

        <div class="md-layout">
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart :type="'consumer'" :num="10000" :size="5000" />
            </div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart
                    :type="'consumer'"
                    :num="100000"
                    :size="5000"
                />
            </div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart
                    :type="'consumer'"
                    :num="1000000"
                    :size="5000"
                />
            </div>
        </div>
        <div class="md-layout">
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            >
                <benchmark-chart
                    :type="'consumer'"
                    :num="10000000"
                    :size="5000"
                />
            </div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            ></div>
            <div
                class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-33"
            ></div>
        </div>
    </div>
</template>

<script>
import BenchmarkChart from "../components/BenchmarkChart.vue";

export default {
    components: {
        BenchmarkChart,
    },
};
</script>
