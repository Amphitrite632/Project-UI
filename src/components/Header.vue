<template>
    <header>
        <a id="headerSiteName" href="/">
            <Logo />
        </a>
        <div id="headerBody">
            <div id="searchBoxContainer">
                <span class="materialIcons" id="searchIcon">search</span>
                <input type="text" autocomplete="off" id="searchBox" placeholder="イラストを検索" />
                <dialog id="searchBoxSuggestDialog">
                    <p id="searchBoxSuggestDialogTip">No suggestions</p>
                    <div class="searchBoxSuggestSection" id="searchResultElementsContainer">{{}}</div>
                </dialog>
            </div>
        </div>
    </header>
</template>

<script lang="ts" setup>
    import Logo from "./Logo.vue"
</script>
