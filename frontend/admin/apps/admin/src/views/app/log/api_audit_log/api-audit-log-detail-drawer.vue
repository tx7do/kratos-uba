<script lang="ts" setup>
import type { auditservicev1_ApiAuditLog } from '#/api';

import { ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Descriptions, DescriptionsItem, Tag } from 'ant-design-vue';

import { successToColor, successToNameWithStatusCode } from '#/api';
import { $t } from '#/locales';

// 与列表 httpMethod 列语义同源：GET 成功绿 / POST 处理蓝 / PUT 警告橙 / DELETE 错误红
const methodColorMap: Record<string, string> = {
  GET: 'success',
  POST: 'processing',
  PUT: 'warning',
  DELETE: 'error',
};

const row = ref<auditservicev1_ApiAuditLog>();

const [Drawer, drawerApi] = useVbenDrawer({
  footer: false,
  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      row.value = drawerApi.getData<{ row: auditservicev1_ApiAuditLog }>()?.row;
    } else {
      row.value = undefined;
    }
  },
});
</script>

<template>
  <Drawer :title="$t('ui.modal.view', { moduleName: $t('page.apiAuditLog.moduleName') })" class="w-full max-w-[800px]">
    <div class="api-audit-detail">
      <Card class="mb-4">
        <template #title>
          <span class="card-title">{{
            $t('page.apiAuditLog.sectionBasic')
          }}</span>
        </template>
        <Descriptions :column="2" bordered>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.createdAt')"
          >{{ formatDateTime(row?.createdAt ?? '') || '-' }}</DescriptionsItem>
          <DescriptionsItem :label="$t('page.apiAuditLog.success')">
            <Tag
              v-if="row?.success !== undefined && row?.success !== null"
              :color="successToColor(row.success)"
            >
              {{ successToNameWithStatusCode(row.success, row?.statusCode ?? 0) }}
            </Tag>
            <span v-else>-</span>
          </DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.username')"
          >{{ row?.username || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.userId')"
          >{{ row?.userId ?? '-' }}</DescriptionsItem>
          <DescriptionsItem :label="$t('page.apiAuditLog.httpMethod')">
            <Tag
              v-if="row?.httpMethod"
              :color="methodColorMap[row.httpMethod] ?? 'default'"
            >
              {{ row.httpMethod }}
            </Tag>
            <span v-else>-</span>
          </DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.statusCode')"
          >{{ row?.statusCode ?? '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.path')"
            :span="2"
          >{{ row?.path || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.requestUri')"
            :span="2"
          >{{ row?.requestUri || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.latencyMs')"
          >{{ row?.latencyMs ?? '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.ipAddress')"
          >{{ row?.ipAddress || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.tenantName')"
          >{{ row?.tenantName || '-' }}</DescriptionsItem>
          <DescriptionsItem :label="$t('page.apiAuditLog.geoLocation')">
            {{ row?.geoLocation?.province || '' }}
            {{ row?.geoLocation?.city || '' }}
          </DescriptionsItem>
        </Descriptions>
      </Card>

      <Card class="mb-4">
        <template #title>
          <span class="card-title">{{
            $t('page.apiAuditLog.sectionDevice')
          }}</span>
        </template>
        <Descriptions :column="2" bordered>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.platform')"
          >{{ row?.deviceInfo?.platform || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.osName')"
          >{{ row?.deviceInfo?.osName || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.browserName')"
          >{{ row?.deviceInfo?.browserName || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.userAgent')"
            :span="2"
          >{{ row?.deviceInfo?.userAgent || '-' }}</DescriptionsItem>
        </Descriptions>
      </Card>

      <Card class="mb-4">
        <template #title>
          <span class="card-title">{{
            $t('page.apiAuditLog.sectionApi')
          }}</span>
        </template>
        <Descriptions :column="2" bordered>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.apiModule')"
          >{{ row?.apiModule || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.apiOperation')"
          >{{ row?.apiOperation || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.apiDescription')"
            :span="2"
          >{{ row?.apiDescription || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.referer')"
            :span="2"
          >{{ row?.referer || '-' }}</DescriptionsItem>
        </Descriptions>
      </Card>

      <Card class="mb-4">
        <template #title>
          <span class="card-title">{{
            $t('page.apiAuditLog.sectionOther')
          }}</span>
        </template>
        <Descriptions :column="2" bordered>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.reason')"
            :span="2"
          >{{ row?.reason || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.spanId')"
          >{{ row?.spanId || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.requestId')"
          >{{ row?.requestId || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.traceId')"
          >{{ row?.traceId || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.logHash')"
            :span="2"
          >{{ row?.logHash || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.signature')"
            :span="2"
          >{{ row?.signature || '-' }}</DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.requestBody')"
            :span="2"
          >
            <pre v-if="row?.requestBody" class="pre-text">{{ row.requestBody }}</pre>
            <span v-else>-</span>
          </DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.requestHeader')"
            :span="2"
          >
            <pre v-if="row?.requestHeader" class="pre-text">{{ row.requestHeader }}</pre>
            <span v-else>-</span>
          </DescriptionsItem>
          <DescriptionsItem
            :label="$t('page.apiAuditLog.response')"
            :span="2"
          >
            <pre v-if="row?.response" class="pre-text">{{ row.response }}</pre>
            <span v-else>-</span>
          </DescriptionsItem>
        </Descriptions>
      </Card>
    </div>
  </Drawer>
</template>

<style scoped>
.api-audit-detail {
  padding-right: 10px;
}

.pre-text {
  margin: 0;
  max-height: 240px;
  overflow: auto;
  white-space: pre-wrap;
  word-break: break-all;
  font-family: inherit;
}
</style>
