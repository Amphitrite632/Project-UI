<script lang="ts" setup>
    import { getOptions } from "../util"
    import Header from "../components/Header.vue"
    import "../css/home.css"
</script>

<script lang="ts">
    export default {
        async mounted() {
            const options = await getOptions()
            document.title = `About - ${options.siteName}`
        },
    }
</script>

<template>
    <Header />
    <div id="body">
        <Illusts />
    </div>
</template>
