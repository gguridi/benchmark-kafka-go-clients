import DashboardLayout from "../pages/Layout/DashboardLayout.vue";
import Producers from "../pages/Producers.vue";
import Consumers from "../pages/Consumers.vue";

const routes = [
    {
        path: "/",
        component: DashboardLayout,
        redirect: "/producers",
        children: [
            {
                path: "producers",
                name: "Producers",
                component: Producers,
            },
            {
                path: "consumers",
                name: "Consumers",
                component: Consumers,
            },
        ],
    },
];

export default routes;
