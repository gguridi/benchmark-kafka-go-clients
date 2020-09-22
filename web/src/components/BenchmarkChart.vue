<template>
    <chart-card
        :chart-data="data"
        :chart-options="options"
        :chart-type="'Bar'"
        :data-background-color="backgroundColor"
    >
        <template slot="content">
            <h4 class="title">{{ winner }}</h4>
            <p class="category">{{ description }}</p>
        </template>
    </chart-card>
</template>
<script>
import { ChartCard } from "@/components";

export default {
    components: {
        ChartCard,
    },
    name: "benchmark-chart",
    props: {
        backgroundColor: {
            type: String,
            default: "",
        },
        type: {
            type: String,
        },
        num: {
            type: Number,
        },
        size: {
            type: Number,
        },
    },
    data() {
        const clients = ["confluent", "sarama", "kafkago"];
        const average = [];
        for (const client of clients) {
            try {
                const data = require(`../results/${this.type}-${client}-${this.num}-${this.size}/results.json`);
                average.push(data["average"]);
            } catch (e) {
                average.push(0);
            }
        }
        const min = Math.min.apply(Math, average.filter(Boolean));
        const max = Math.max.apply(Math, average);
        const winner = clients[average.findIndex((value) => value === min)];
        const high = Math.round(max) + 1;

        return {
            winner: winner,
            description: `${this.num} messages/${this.size} bytes each`,
            data: {
                labels: clients,
                series: [average],
            },
            options: {
                axisX: {
                    showGrid: false,
                },
                low: 0,
                high: max,
                chartPadding: {
                    top: 0,
                    right: 0,
                    bottom: 0,
                    left: 0,
                },
            },
        };
    },
};
</script>
