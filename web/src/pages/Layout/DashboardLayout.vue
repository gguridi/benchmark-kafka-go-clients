<template>
    <div class="wrapper" :class="{ 'nav-open': $sidebar.showSidebar }">
        <notifications></notifications>

        <side-bar
            :sidebar-item-color="sidebarBackground"
            :sidebar-background-image="sidebarBackgroundImage"
            :title="'Kafka Go Clients'"
        >
            <mobile-menu slot="content"></mobile-menu>
            <sidebar-link to="/producers">
                <md-icon>dashboard</md-icon>
                <p>Producers</p>
            </sidebar-link>
            <sidebar-link to="/consumers">
                <md-icon>dashboard</md-icon>
                <p>Consumers</p>
            </sidebar-link>
        </side-bar>

        <div class="main-panel">
            <dashboard-content></dashboard-content>
            <content-footer v-if="!$route.meta.hideFooter"></content-footer>
        </div>
    </div>
</template>

<script>
import ContentFooter from "./ContentFooter.vue";
import DashboardContent from "@/pages/Layout/Content.vue";
import MobileMenu from "@/pages/Layout/MobileMenu.vue";

export default {
    components: {
        DashboardContent,
        ContentFooter,
        MobileMenu,
    },
    data() {
        return {
            sidebarBackground: "green",
            sidebarBackgroundImage: require("@/assets/img/sidebar-2.jpg"),
        };
    },
};
</script>
