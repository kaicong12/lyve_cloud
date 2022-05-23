import React, { useEffect, useState } from "react";
import { Card, Row, Space, Table, Pagination, Button, Typography } from "antd";
// import { useDispatch, useSelector } from '@@/plugin-dva/exports';
import { useNavigate } from "react-router-dom";

// import type { ConnectState } from '@/models/connect';
import TableSearchFilter, { StatusOption } from "../utils/TableSearchFilter";
import { MigrationsColumnFormatter } from "../utils/MigrationsColumnFormatter";

export const MigrationList: React.FC = () => {
  // const dispatch = useDispatch();
  const navigate = useNavigate();
  // Migration listing states
  const [pageConfig, setPageConfig] = useState({
    current: 1,
    pageSize: 5,
  });
  const [filterInfo, setFilterInfo] = useState({});
  const [sortInfo, setSortInfo] = useState({ order: null, ascending: true });
  // const modelList = useSelector((state: ConnectState) => state.activeLearning.modelList);

  useEffect(() => {
    // dispatch({
    //   type: 'activeLearning/fetchALModelList',
    //   payload: {
    //     filter: filterInfo,
    //     order: sortInfo.order,
    //     ascending: sortInfo.ascending,
    //     pageSize: pageConfig.pageSize,
    //     page: pageConfig.current,
    //   },
    // });
  }, [filterInfo, pageConfig, sortInfo]);

  const handleChangeMigration = (_: any, __: any, sorter: any) => {
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

  const handleLinkClick = (value: number) => {
    navigate(`/lyve_cloud_migration/${value}`);
  };

  return (
    <div style={{ width: "95vh" }}>
      <Space style={{ width: "100%" }} direction="vertical">
        <Button
          style={{
            margin: "5px",
            borderRadius: "10px",
          }}
          type="primary"
          size="middle"
          onClick={() => {
            navigate("/lyve_cloud_migration/create_migration");
          }}
        >
          + New Migration
        </Button>
        <Card style={{ borderRadius: "15px" }}>
          <Space style={{ width: "100%" }} direction="vertical">
            <Row justify="space-between">
              <Typography.Title level={4}>Migration History</Typography.Title>
              <TableSearchFilter
                columns={MigrationsColumnFormatter({ handleLinkClick })}
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
              columns={MigrationsColumnFormatter({ handleLinkClick })}
              // dataSource={modelList.data}
              onChange={handleChangeMigration}
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
      </Space>
    </div>
  );
};

export default MigrationList;
