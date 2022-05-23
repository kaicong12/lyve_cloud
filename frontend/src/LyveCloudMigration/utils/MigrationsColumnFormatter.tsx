import { Badge, Typography } from "antd";
import React from "react";

enum MigrationsColumnIndex {
  NAME = "name",
  STATUS = "status",
  NO_OF_MIGRATED_OBJS = "no_migrat_objs",
  NO_OF_FAILED_OBJS = "no_failed_objs",
  EXECUTION_TIME = "exec_time",
}

enum ActivitiesColumnIndex {
  OBJECT = "object",
  STATUS = "status",
  OBJECT_SIZE = "obj_size",
  EXECUTION_TIME = "exec_time",
}

type MigrationsColumnFormatterProps = {
  handleLinkClick: any;
};

enum BadgeComponentStatus {
  PROCESSING = "processing",
  ERROR = "error",
  WARNING = "warning",
  SUCCESS = "success",
}

enum BadgeComponentStatusText {
  PROCESSING = "In Progress",
  ERROR = "Error",
  DONE = "Done",
  NOT_STARTED = "Not started",
}

const convertToBadgeStatusAndText: (
  migrationStatus: "in_progress" | "done" | "error" | undefined
) => {
  status: BadgeComponentStatus;
  text: BadgeComponentStatusText;
} = (migrationStatus) => {
  const toRet: {
    status: BadgeComponentStatus;
    text: BadgeComponentStatusText;
  } = {
    status: BadgeComponentStatus.ERROR,
    text: BadgeComponentStatusText.ERROR,
  };
  switch (migrationStatus) {
    case "in_progress":
      toRet.status = BadgeComponentStatus.PROCESSING;
      toRet.text = BadgeComponentStatusText.PROCESSING;
      break;
    case "done":
      toRet.status = BadgeComponentStatus.SUCCESS;
      toRet.text = BadgeComponentStatusText.DONE;
      break;
    case "error":
      toRet.status = BadgeComponentStatus.ERROR;
      toRet.text = BadgeComponentStatusText.ERROR;
      break;
    default:
      break;
  }
  return toRet;
};

export const MigrationsColumnFormatter = (
  props: MigrationsColumnFormatterProps
) => {
  const { handleLinkClick } = props;

  return [
    {
      title: "Name",
      dataIndex: MigrationsColumnIndex.NAME,
      key: MigrationsColumnIndex.NAME,
      sorter: true,
      searchFilter: true,
      render: (text: any, record: any) => (
        <Typography.Link onClick={() => handleLinkClick(record.id)}>
          {text}
        </Typography.Link>
      ),
    },
    {
      title: "Status",
      dataIndex: MigrationsColumnIndex.STATUS,
      key: MigrationsColumnIndex.STATUS,
      searchFilter: true,
      render: (obj: any) => {
        const { status: badgeStatus, text: badgeText } =
          convertToBadgeStatusAndText(obj);
        return <Badge status={badgeStatus} text={badgeText} />;
      },
    },
    {
      title: "Number of migrated objects",
      dataIndex: MigrationsColumnIndex.NO_OF_MIGRATED_OBJS,
      sorter: true,
    },
    {
      title: "Number of failed objects",
      dataIndex: MigrationsColumnIndex.NO_OF_FAILED_OBJS,
      sorter: true,
    },
    {
      title: "Execution time",
      dataIndex: MigrationsColumnIndex.EXECUTION_TIME,
      key: MigrationsColumnIndex.EXECUTION_TIME,
      render: (obj: any) => new Date(obj).toLocaleDateString(),
    },
  ];
};

export const ActivitiesColumnFormatter = () => {
  return [
    {
      title: "Object",
      dataIndex: ActivitiesColumnIndex.OBJECT,
      key: ActivitiesColumnIndex.OBJECT,
      sorter: true,
      searchFilter: true,
    },
    {
      title: "Status",
      dataIndex: ActivitiesColumnIndex.STATUS,
      key: ActivitiesColumnIndex.STATUS,
      searchFilter: true,
      render: (obj: any) => {
        const { status: badgeStatus, text: badgeText } =
          convertToBadgeStatusAndText(obj);
        return <Badge status={badgeStatus} text={badgeText} />;
      },
    },
    {
      title: "Object size",
      dataIndex: ActivitiesColumnIndex.OBJECT_SIZE,
      sorter: true,
    },
    {
      title: "Execution time",
      dataIndex: ActivitiesColumnIndex.EXECUTION_TIME,
      key: ActivitiesColumnIndex.EXECUTION_TIME,
      render: (obj: any) => new Date(obj).toLocaleDateString(),
    },
  ];
};
