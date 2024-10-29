<template>
    <ul>
        <li v-for="(node, index) in data" :key="index">
            {{ node?.name }}
            <Tree :data="(node.children as any)"></Tree>
        </li>
    </ul>
</template>
<script setup lang="ts">
import { Node } from "../types/Node";

interface Props {
    data: Node[]
}

let { data } = defineProps<Props>()

// let { data } = withDefaults(defineProps<Props>(), {
//     data: [] as any,
// })

</script>
<style scoped lang="scss"></style>