<template>
    <div class="illustContainer" v-for="illust in illusts">
        <div class="illust" v-bind:style="`background-image: url(/asset/illust/${illust.illustID}.webp); --illust-height-ratio: ${illust.heightRatio}`"></div>
        <a class="illustOriginalPostAnchor" target="_blank" v-bind:href="`https://${illust.originalURL}`" :data-is-original-post-exists="`${illust.originalURL != null}`">オリジナルの投稿</a>
        <div class="illustTags">
            <a class="illustTag" v-for="tag in illust.tags" v-bind:href="`/?tagID=${tag.tagID}`" :data-tag-type="tag.type">{{ tag.officialName }}</a>
        </div>
        <div class="illustInfoEditButton" v-on:click="editIllustInfo(illust.illustID)">登録情報を編集</div>
    </div>
</template>

<script lang="ts">
    import { IllustInfo } from "../v13s.types"

    function activateIllustInfoEditDialog(illustID: string) {
        history.pushState("", "", "/edit/illust/")
        window.localStorage.setItem("editIllustTargetID", illustID)
        const illustInfoEditDialog = document.getElementById("illustInfoEditDialog") as HTMLDialogElement
        const bodyFilter = document.getElementById("bodyFilter") as HTMLDivElement

        illustInfoEditDialog.open = true
        bodyFilter.style.animationName = "activateBodyFilter"
        illustInfoEditDialog.style.animationName = "activateModal"
        //TODO: ダイアログへの値の受け渡しを書く
    }

    function disableIllustInfoEditDialog() {
        history.pushState("", "", "/")

        const bodyFilter = document.getElementById("bodyFilter") as HTMLDivElement
        const illustInfoEditDialog = document.getElementById("illustInfoEditDialog") as HTMLDivElement
        illustInfoEditDialog.style.animationName = "disableModal"
        bodyFilter.style.animationName = "disableBodyFilter"
    }

    window.addEventListener("popstate", () => {
        if ((window.location.pathname = "/edit/illust/")) {
            const illustID = window.localStorage.getItem("editIllustTargetID")
            if (illustID != null) {
                activateIllustInfoEditDialog(illustID)
            } else {
                disableIllustInfoEditDialog()
            }
        } else {
            disableIllustInfoEditDialog()
        }
    })

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

        computed: {
            editIllustInfo() {
                return activateIllustInfoEditDialog
            },
        },
    }
</script>
