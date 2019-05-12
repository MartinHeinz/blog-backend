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
                    v-on:click="toggle_social_sharing"
                    slot-scope="{ hover }"
                    :class="`${hover? 'icon-active': 'icon-inactive'}`">fas fa-share-alt</v-icon>
            </v-hover>
        </li>
      </ul>

      <span class="info" v-show="prev_active">Previous post</span>
      <span class="info" v-show="next_active">Next post</span>
      <span class="info" v-show="top_active">Back to top</span>
      <span class="info" v-show="share_active">Share post</span>
    </span>
</template>

<script>

export default {
    name: 'MenuActions',
    props: {
        url_previous: String,
        url_next: String,
    },
    methods: {
        toggle_social_sharing() {
            this.$root.$emit('toggle_social_sharing');
        },
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
    #actions {
        float: right;
        margin-top: 2rem;
        margin-right: 2rem;
        width: 100%;
        text-align: right;
    }
    #actions ul {
        display: block;
    }
    #actions .info {
        font-style: italic;
    }

    .icon-inactive {
        color: #c9cacc;
    }
    .icon-active {
        color: #d480aa;
        cursor: pointer;
    }
</style>
