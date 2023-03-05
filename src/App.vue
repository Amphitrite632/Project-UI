<template>
    <router-view></router-view>
</template>

<script lang="ts" setup>
    const comment = document.createComment("\n        __   _____      \n       /  | |____ |     \n__   __`| |     / / ___ \n\\ \\ / / | |     \\ \\/ __|\n \\ V / _| |_.___/ /\\__ \\\n  \\_/  \\___/\\____/ |___/\n                        \n-- Vivliothikarios(v13s) by v13s team --\n\nHello there! Thank you for using v13s.\nWe are always looking for curious developer... like you!\ngithub.com/Vivliothikarios\n\n")

    document.documentElement.prepend(comment)
</script>
