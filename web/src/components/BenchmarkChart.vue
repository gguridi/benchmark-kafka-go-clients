<template>
    <chart-card
        :chart-data="data"
        :chart-options="options"
        :chart-type="'Bar'"
        :data-background-color="background"
    >
        <template slot="content">
            <h4 class="title">Winner: {{ winner }}</h4>
            <p class="category">Benchmark: {{ description }}</p>
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
        const backgrounds = ["purple", "blue", "green", "orange", "red"];
        const background =
            backgrounds[Math.floor(Math.random() * backgrounds.length)];
        const producers = ["confluent", "sarama", "kafkago"];
        const consumers = [
            "confluent-poll",
            "confluent-channel",
            "sarama",
            "kafkago",
        ];
        const clients = this.type == "producer" ? producers : consumers;
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
        const low = Math.round(min) - 1;
        const high = Math.round(max) + 1;

        return {
            winner: winner,
            description: `${this.num} messages/${this.size} bytes each`,
            background: background,
            data: {
                labels: clients,
                series: [average],
            },
            options: {
                axisX: {
                    showGrid: false,
                },
                low: low > 0 ? low : 0,
                high: high,
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
