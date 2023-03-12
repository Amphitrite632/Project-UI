<script lang="ts" setup>
    import { editIllustInfo, disableIllustInfoEditDialog } from "../util"
    import Header from "../components/Header.vue"
    import Illusts from "../components/Illusts.vue"
    import IllustInfoEditDialog from "../components/IllustInfoEditDialog.vue"
    import Footer from "../components/Footer.vue"
    import BodyFilter from "../components/BodyFilter.vue"
    import "../css/home.css"
    import { IllustInfo } from "../v13s.types"
    
    document.title = `Home - ${document.title}`
</script>

<script lang="ts">
    window.addEventListener("popstate", () => {
        if (window.location.pathname != "/edit/illust/") {
            disableIllustInfoEditDialog()
        }
    })

    export default {
        data: () => {
            const illustInfo: IllustInfo = {
                illustID: "",
                hash: "",
                originalURL: "",
                heightRatio: "",
                tags: [],
                createdAt: "",
            }
            return {
                illustInfo: illustInfo,
            }
        },

        methods: {
            async onillustInfoEditButtonClick(illust: IllustInfo) {
                this.illustInfo = illust
                editIllustInfo()
            },
        },
    }
</script>

<template>
    <Header />
    <div id="body">
        <Illusts v-on:illustInfoEditButtonClick="onillustInfoEditButtonClick" />
    </div>
    <KeepAlive>
        <IllustInfoEditDialog v-model:illust-info="illustInfo" />
    </KeepAlive>
    <Footer />
    <BodyFilter />
</template>
