interface NotificationItem {
  id: number;
  messageId: number;
  date: string;
  isRead?: boolean;
  message: string;
  title: string;
}

export type { NotificationItem };
