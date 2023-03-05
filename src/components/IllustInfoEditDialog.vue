<template>
    <dialog id="illustInfoEditDialog" class="centerDialog aboveBodyFilter">
        <p>{{ illustInfo.illustID }}</p>
    </dialog>
</template>

<script lang="ts">
    import { PropType } from "vue"
    import { IllustInfo } from "../v13s.types"

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

        props: {
            illustInfo: {
                type: Object as PropType<IllustInfo>,
            },
        },
    }
</script>
