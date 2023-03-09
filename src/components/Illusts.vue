<template>
    <div class="illustContainer" v-for="illust in illusts">
        <div class="illust" v-bind:style="`background-image: url(/asset/illust/${illust.illustID}.webp); --illust-height-ratio: ${illust.heightRatio}`"></div>
        <a class="illustOriginalPostAnchor" target="_blank" v-bind:href="`https://${illust.originalURL}`" :data-is-original-post-exists="`${illust.originalURL != null}`">オリジナルの投稿</a>
        <div class="illustTags">
            <a class="illustTag" v-for="tag in illust.tags" v-bind:href="`/?tagID=${tag.tagID}`" :data-tag-type="tag.type">{{ tag.officialName }}</a>
        </div>
        <div class="illustInfoEditButton" @click="callIllustInfoEditHandler(illust)">登録情報を編集</div>
    </div>
</template>

<script lang="ts">
    import { disableIllustInfoEditDialog } from "../util"
    import { IllustInfo } from "../v13s.types"

    export default {
        data: () => {
            const illusts: IllustInfo[] = []
            return {
                illusts: illusts,
            }
        },

        async mounted() {
            const illustJSON = await fetch("/api/illusts/")
            this.illusts = await illustJSON.json()
            window.addEventListener("bodyFilterClick", disableIllustInfoEditDialog)
        },
        methods: {
            callIllustInfoEditHandler(illust: IllustInfo) {
                this.$emit("illustInfoEditButtonClick", illust)
            },
        },
    }
</script>
