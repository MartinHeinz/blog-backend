<template>
    <span id="actions">
      <ul>
        <li><v-icon :href="url_previous">fas fa-chevron-left</v-icon></li>
        <li><v-icon :href="url_next">fas fa-chevron-right</v-icon></li>
        <li><v-icon :large=false href="#">fas fa-chevron-up</v-icon></li>
        <li><v-icon href="#">fas fa-share-alt</v-icon></li>
      </ul>
      <span id="i-prev" class="info" style="display: none;">Previous post</span>
      <span id="i-next" class="info" style="display: none;">Next post</span>
      <span id="i-top" class="info" style="display: none;">Back to top</span>
      <span id="i-share" class="info" style="display: none;">Share post</span>
    </span>
</template>

<script>

export default {
    name: 'MenuActions',
    props: {
        url_previous: String,
        url_next: String,
    },
};
</script>
