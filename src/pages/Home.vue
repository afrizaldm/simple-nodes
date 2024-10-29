<template>
    <div>
        <Tree :data="(nodes as any)"></Tree>
    </div>
</template>
<script setup lang="ts">
import { onMounted } from 'vue';

import { nodes, load } from "../services/node";

import Tree from "../components/Tree.vue";

onMounted(async () => {
    load()
})

</script>
<style scoped lang="scss"></style>