<script lang="ts" setup>
    import Header from "../components/Header.vue"
    import "../css/home.css"

    document.title = `About - ${document.title}`
</script>

<template>
    <Header />
    <div id="body">
        <Illusts />
    </div>
</template>
