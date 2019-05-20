<template>
    <div>
        <a :href="href">
        <div v-bind:class="setStyle">{{ label }}</div>
        <TableOfContents
            v-for="(node, index) in nodes"
            :nodes="node.nodes"
            :key="index"
            :depth="depth + 1"
            :label="node.label"
            :href="node.href">
        </TableOfContents>
        </a>
    </div>
</template>

<script>

export default {
    name: 'TableOfContents',
    props: ['label', 'nodes', 'depth', 'href'],
    computed: {
        setStyle() {
            return `toc-level-${this.depth}`;
        },
    },
};
</script>

<style>

</style>
