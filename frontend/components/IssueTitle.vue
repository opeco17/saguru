<template>
    <a 
        :href="url"
        target="_blank"
        class="issue-title mr-2 text-subtitle-1"
    >
        <slot></slot>
    </a>
</template>

<script>
export default {
    props: {
        url: String
    }
}
</script>

<style scoped>
.issue-title {
    text-decoration: none;
    color: black;
    transition: color 0.1s;
}
.issue-title:hover {
    color: purple;
}
</style>