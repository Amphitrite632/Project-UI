<template>
    <dialog id="illustInfoEditDialog" class="centerDialog illustInfoDialog aboveBodyFilter">
        <div id="illustInfoEditPreviewImage">
            <img v-bind:src="`/asset/illust/${illustInfo.illustID}.webp`" />
        </div>
        <div id="illustInfoEditInputField">
            <div data-position="upper">
                <p>登録情報を編集</p>
                <div class="illustInfoInputContainer">
                    <label for="illustOriginalURLInput">オリジナルの投稿のURL</label>
                    <input id="illustOriginalURLInput" />
                    <p data-info="icon">URLが空の場合、登録情報は「リンク切れ」になります</p>
                </div>
                <div class="illustInfoInputContainer">
                    <label for="illustTagAddInput">タグを追加</label>
                    <input id="illustTagAddInput" />
                    <div id="illustTagSuggestionContainer">
                        <p>No suggestions</p>
                        <div v-for="suggestion in suggestions" class="illustTagSuggestion" @click="addTagToIllustInfo(suggestion.tagID)">
                            {{ suggestion.officialName }}
                        </div>
                    </div>
                    <p></p>
                </div>
            </div>
            <div data-position="lower">
                <div class="illustInfoEditTag" v-for="tag in (tags = illustInfo.tags)" :data-tag-id="tag.tagID" :data-tag-type="tag.type" data-is-exists="true">
                    {{ tag.officialName }}<span class="tagStatusToggleButton" @click="(event) => toggleTagStatus(event.target as HTMLSpanElement)"><CrossIcon style="pointer-events: none" /></span>
                </div>
            </div>
        </div>
    </dialog>
</template>

<script lang="ts" setup>
    import CrossIcon from "./icon/Cross.vue"
</script>

<script lang="ts">
    import { PropType } from "vue"
    import { IllustInfo, TagInfo } from "../v13s.types"

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
            const suggestions: TagInfo[] = []
            return {
                illustInfo: illustInfo,
                tags: illustInfo.tags,
                suggestions: suggestions,
            }
        },

        props: {
            illustInfo: {
                type: Object as PropType<IllustInfo>,
            },
        },

        methods: {
            addTagToIllustInfo(tagID: string) {},
            toggleTagStatus(button: HTMLSpanElement) {
                const tagElement = button.parentElement as HTMLDivElement
                if ((this.$props.illustInfo as IllustInfo).tags.find((tag) => tag.tagID == tagElement.dataset.tagId) != null) {
                    if (tagElement.dataset.isExists == "true") {
                        button.style.transform = "rotate(45deg)"
                        tagElement.dataset.isExists = "false"
                    } else {
                        button.style.transform = "rotate(0deg)"
                        tagElement.dataset.isExists = "true"
                    }
                }
            },
        },
    }
</script>
