<script lang="ts" setup>
import type {
  ubaservicev1_IDMapping,
  ubaservicev1_UserTag as UserTag,
} from '#/generated/api/admin/service/v1';

import { reactive, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { LucideSearch } from '@vben/icons';

import {
  Card,
  InputNumber,
  message,
  TabPane,
  Tabs,
  Tag,
  Timeline,
  TimelineItem,
} from 'ant-design-vue';

import {
  fetchListIDMappings,
  fetchListUserTags,
  PaginationQuery,
  useBehaviorSequence,
} from '#/api';
import { $t } from '#/locales';

const userId = ref<number>();
const appId = ref<number>();

// 行为序列：用户 ID 就绪后自动查询
const sequenceReq = reactive<Record<string, any>>({
  userId: 0,
  limit: 100,
  timeRange: {
    startMs: Date.now() - 30 * 86_400_000,
    endMs: Date.now(),
  },
});

const { data: sequence, isFetching: sequenceLoading } = useBehaviorSequence(
  sequenceReq as any,
  {
    queryKey: ['userBehaviorSequence', sequenceReq],
    enabled: () => sequenceReq.userId > 0,
  } as any,
);

const tags = ref<UserTag[]>([]);
const idMappings = ref<ubaservicev1_IDMapping[]>([]);

async function handleSearch() {
  if (!userId.value || userId.value <= 0) {
    message.warning($t('page.userProfile.userIdPlaceholder'));
    return;
  }

  sequenceReq.userId = userId.value;
  if (appId.value && appId.value > 0) {
    sequenceReq.appId = appId.value;
  }

  const [tagResp, mappingResp] = await Promise.allSettled([
    fetchListUserTags(
      new PaginationQuery({
        paging: { page: 1, pageSize: 100 },
        formValues: { userId: userId.value },
      }),
    ),
    fetchListIDMappings(
      new PaginationQuery({
        paging: { page: 1, pageSize: 100 },
        formValues: { idValue: String(userId.value) },
      }),
    ),
  ]);

  if (tagResp.status === 'fulfilled') {
    tags.value = (tagResp.value.items || []) as UserTag[];
  }
  if (mappingResp.status === 'fulfilled') {
    idMappings.value = (mappingResp.value.items ||
      []) as ubaservicev1_IDMapping[];
  }
}

function fmtTime(ts?: Date | number | string) {
  if (!ts) return '-';
  const d = typeof ts === 'object' ? (ts as Date) : new Date(Number(ts));
  if (Number.isNaN(d.getTime())) return String(ts);
  return d.toLocaleString();
}
</script>

<template>
  <Page auto-content-height>
    <Card size="small" class="mb-4">
      <div class="flex flex-wrap items-center gap-3">
        <span>{{ $t('page.userProfile.userId') }}</span>
        <InputNumber
          v-model:value="userId"
          :min="1"
          :precision="0"
          :placeholder="$t('page.userProfile.userIdPlaceholder')"
          style="width: 180px"
          @press-enter="handleSearch"
        />
        <span>{{ $t('page.userProfile.appId') }}</span>
        <InputNumber
          v-model:value="appId"
          :min="1"
          :precision="0"
          :placeholder="$t('page.userProfile.appIdOptional')"
          style="width: 160px"
        />
        <a-button type="primary" @click="handleSearch">
          <template #icon>
            <LucideSearch />
          </template>
          {{ $t('page.userProfile.search') }}
        </a-button>
      </div>
    </Card>

    <Card v-if="sequenceReq.userId > 0" size="small">
      <Tabs>
        <TabPane key="sequence" :tab="$t('page.userProfile.tab.sequence')">
          <a-spin :spinning="sequenceLoading">
            <Timeline v-if="(sequence?.events || []).length > 0" class="mt-2">
              <TimelineItem
                v-for="(event, index) in sequence?.events"
                :key="index"
              >
                <template #label>
                  <span class="text-gray-400">
                    {{ fmtTime(event.timestamp) }}
                  </span>
                </template>
                <div>
                  <Tag color="blue">{{ event.eventName }}</Tag>
                  <span v-if="event.referer" class="text-sm">
                    {{ event.referer }}
                  </span>
                </div>
              </TimelineItem>
            </Timeline>
            <a-empty v-else :description="$t('page.userProfile.noEvents')" />
          </a-spin>
        </TabPane>

        <TabPane key="tags" :tab="$t('page.userProfile.tab.tags')">
          <a-table
            :data-source="tags"
            :pagination="{ pageSize: 10 }"
            row-key="id"
            size="small"
          >
            <a-table-column
              :title="$t('page.userProfile.tag.tagKey')"
              data-index="tagKey"
              :width="160"
            />
            <a-table-column
              :title="$t('page.userProfile.tag.tagValue')"
              data-index="tagValue"
              :width="180"
            />
            <a-table-column
              :title="$t('page.userProfile.tag.source')"
              data-index="source"
              :width="110"
            >
              <template #default="{ text }">
                <Tag>{{ text || '-' }}</Tag>
              </template>
            </a-table-column>
            <a-table-column
              :title="$t('page.userProfile.tag.confidence')"
              data-index="confidence"
              :width="100"
            />
            <a-table-column
              :title="$t('page.userProfile.tag.effectiveTime')"
              data-index="effectiveTime"
              :width="170"
            >
              <template #default="{ text }">
                {{ fmtTime(text) }}
              </template>
            </a-table-column>
          </a-table>
        </TabPane>

        <TabPane key="ids" :tab="$t('page.userProfile.tab.ids')">
          <a-table
            :data-source="idMappings"
            :pagination="{ pageSize: 10 }"
            row-key="id"
            size="small"
          >
            <a-table-column
              :title="$t('page.userProfile.id.idType')"
              data-index="idType"
              :width="110"
            />
            <a-table-column
              :title="$t('page.userProfile.id.idValue')"
              data-index="idValue"
              :width="220"
            />
            <a-table-column
              :title="$t('page.userProfile.id.globalUserId')"
              data-index="globalUserId"
              :width="200"
            />
            <a-table-column
              :title="$t('page.userProfile.id.confidence')"
              data-index="confidence"
              :width="100"
            />
            <a-table-column
              :title="$t('page.userProfile.id.firstSeen')"
              data-index="firstSeen"
              :width="170"
            >
              <template #default="{ text }">
                {{ fmtTime(text) }}
              </template>
            </a-table-column>
          </a-table>
        </TabPane>
      </Tabs>
    </Card>

    <Card v-else size="small">
      <a-empty :description="$t('page.userProfile.empty')" />
    </Card>
  </Page>
</template>

<style scoped>
:deep(.ant-descriptions-item-label) {
  min-width: 120px;
}
</style>
