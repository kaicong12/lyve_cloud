import React, { useEffect, useState } from "react";
import {
  Button,
  Card,
  Col,
  Descriptions,
  Pagination,
  Progress,
  Row,
  Space,
  Table,
  Timeline,
  Typography,
} from "antd";
import { useNavigate, useParams } from "react-router-dom";
import { ArrowRightOutlined, LeftOutlined } from "@ant-design/icons";

import TableSearchFilter, { StatusOption } from "../utils/TableSearchFilter";
import { ActivitiesColumnFormatter } from "../utils/MigrationsColumnFormatter";
import "./index.css";

const MigrationDetails: React.FC = () => {
  const { migrationId } = useParams();
  console.log(migrationId);
  const navigate = useNavigate();
  // Activity listing states
  const [pageConfig, setPageConfig] = useState({
    current: 1,
    pageSize: 5,
  });
  const [filterInfo, setFilterInfo] = useState({});
  const [sortInfo, setSortInfo] = useState({ order: null, ascending: true });

  const handleChangeActivity = (_: any, __: any, sorter: any) => {
    if (Object.keys(sorter).length === 0) {
      return;
    }

    if (!sorter.order && sortInfo.order) {
      setSortInfo({
        ...sortInfo,
        order: null,
      });
    } else if (
      sorter.field !== sortInfo.order ||
      sorter.order !== sortInfo.ascending
    ) {
      setSortInfo({
        order: sorter.field,
        ascending: sorter.order === "ascend",
      });
    }
  };

  const handlePaginationChange = (page: number, pageSize?: number) => {
    setPageConfig({
      pageSize: pageSize || pageConfig.pageSize,
      current: page,
    });
  };

  const handleSearchFilter = (value: any, column: any) => {
    const filter = {};
    // filter[column] = value;
    setFilterInfo(filter);
  };

  const handleReset = () => {
    setFilterInfo({});
  };

  const renderMigrationInfo = () => {
    return (
      <Card className="cardRender">
        <Row>
          <Typography.Title level={4}>migrating...</Typography.Title>
        </Row>
        <Row wrap={false}>
          <Descriptions>
            <Descriptions.Item label="Status">tbc...</Descriptions.Item>
            <Descriptions.Item label="AWS bucket">tbc...</Descriptions.Item>
            <Descriptions.Item label="Created at">tbc...</Descriptions.Item>
          </Descriptions>
        </Row>
      </Card>
    );
  };

  const renderMigrationProgress = () => {
    return (
      <Card className="cardRender">
        <Row>
          <Typography.Title level={4}>Progress</Typography.Title>
        </Row>
        <Col span={10}>
          <Progress percent={80} />
        </Col>
      </Card>
    );
  };

  const renderActivitiesTable = () => {
    return (
      <Card className="cardRender">
        <Space style={{ width: "100%" }} direction="vertical">
          <Row justify="space-between">
            <Typography.Title level={4}>Migration History</Typography.Title>
            <TableSearchFilter
              columns={ActivitiesColumnFormatter()}
              handleSearchFilter={handleSearchFilter}
              handleReset={handleReset}
              options={[StatusOption.IN_PROGRESS, StatusOption.DONE]}
            />
          </Row>
          <Table
            locale={{
              sortTitle: "Sort",
              triggerDesc: "Click sort by descend",
              triggerAsc: "Click sort by ascend",
              cancelSort: "Click to cancel sort",
            }}
            columns={ActivitiesColumnFormatter()}
            // dataSource={modelList.data}
            onChange={handleChangeActivity}
            pagination={false}
            // loading={modelList.loading}
          />
          <Row justify="end">
            <Pagination
              showSizeChanger
              onChange={handlePaginationChange}
              pageSize={pageConfig.pageSize}
              current={pageConfig.current}
              // total={modelList.data?.length}
              showTotal={(total, range) =>
                `${range[0]}-${range[1]} of ${total} items`
              }
            />
          </Row>
        </Space>
      </Card>
    );
  };

  const renderErrorLog = () => {
    const logs = [
      { time: "2022", log: "test" },
      { time: "2021", log: "try" },
    ];
    const timelineItems = logs.map((log) => {
      return (
        <Timeline.Item label={log.time} color="#4615B2">
          {log.log}
        </Timeline.Item>
      );
    });

    return (
      <Card className="cardRender" style={{ marginBottom: "25px" }}>
        <Space style={{ width: "100%" }} direction="vertical">
          <Row>
            <Typography.Title level={4}>Error log</Typography.Title>
          </Row>
          <Timeline
            mode="left"
            style={{ fontFamily: "monospace", fontWeight: "bold" }}
          >
            <Timeline.Item label="Timestamp" dot={<ArrowRightOutlined />}>
              Message
            </Timeline.Item>
            {timelineItems}
          </Timeline>
        </Space>
      </Card>
    );
  };

  return (
    <div style={{ width: "95vh" }}>
      <Space direction="vertical" size="middle" style={{ width: "100%" }}>
        <Row style={{ justifyContent: "space-between" }}>
          <Button
            style={{
              borderRadius: "10px",
            }}
            onClick={() => navigate("/lyve_cloud_migration")}
          >
            <LeftOutlined /> Back
          </Button>
          <Button
            style={{
              borderRadius: "10px",
            }}
            type="primary"
            danger
          >
            Terminate Job
          </Button>
        </Row>
        {renderMigrationInfo()}
        {renderMigrationProgress()}
        {renderActivitiesTable()}
        {renderErrorLog()}
      </Space>
    </div>
  );
};

export default MigrationDetails;
