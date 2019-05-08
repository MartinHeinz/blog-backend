<template>
    <span id="actions">
      <ul>
        <li>
            <v-hover>
            <v-icon :href="url_previous" @mouseover="prev_active = true" @mouseout="prev_active = false"
                    slot-scope="{ hover }"
                    :class="`${hover? 'icon-active': 'icon-inactive'}`">fas fa-chevron-left</v-icon>
            </v-hover>
        </li>
        <li>
            <v-hover>
            <v-icon :href="url_next" @mouseover="next_active = true" @mouseout="next_active = false"
                    slot-scope="{ hover }"
                    :class="`${hover? 'icon-active': 'icon-inactive'}`">fas fa-chevron-right</v-icon>
            </v-hover>
        </li>
        <li>
            <v-hover>
            <v-icon :large=false href="#" @mouseover="top_active = true" @mouseout="top_active = false"
                    slot-scope="{ hover }"
                    :class="`${hover? 'icon-active': 'icon-inactive'}`">fas fa-chevron-up</v-icon>
            </v-hover>
        </li>
        <li>
            <v-hover>
            <v-icon href="#" @mouseover="share_active = true" @mouseout="share_active = false"
                    slot-scope="{ hover }"
                    :class="`${hover? 'icon-active': 'icon-inactive'}`">fas fa-share-alt</v-icon>
            </v-hover>
        </li>
      </ul>

        <span id="i-prev" class="info" v-show="prev_active">Previous post</span>
        <span id="i-next" class="info" v-show="next_active">Next post</span>
      <span id="i-top" class="info" v-show="top_active">Back to top</span>
      <span id="i-share" class="info" v-show="share_active">Share post</span>
    </span>
</template>

<script>

export default {
    name: 'MenuActions',
    props: {
        url_previous: String,
        url_next: String,
    },
    data() {
        return {
            prev_active: false,
            next_active: false,
            top_active: false,
            share_active: false,
        };
    },
};
</script>

<style>
    .icon-inactive {
        color: #c9cacc;
    }
    .icon-active {
        color: #d480aa;
        cursor: pointer;
    }
</style>
