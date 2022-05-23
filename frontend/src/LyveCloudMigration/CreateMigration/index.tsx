import React, { useState } from "react";
import {
  Button,
  Card,
  Col,
  DatePicker,
  Form,
  InputNumber,
  Row,
  Typography,
} from "antd";
import { CaretRightOutlined } from "@ant-design/icons";
import { renderAmazonS3FormItems, renderLyveS3FormItems } from "./utils";

const CreateMigration: React.FC = () => {
  const [form] = Form.useForm();
  const [showFilter, setShowFilter] = useState(false);

  return (
    <div style={{ width: "95vh" }}>
      <Card style={{ borderRadius: "15px" }}>
        <Row justify="center">
          <Col span={18}>
            <Form
              form={form}
              labelCol={{ span: 6 }}
              wrapperCol={{ span: 18 }}
              onSubmitCapture={() => console.log(form.getFieldsValue())}
            >
              <Col offset={6}>
                <Typography.Title level={4}>
                  AWS S3 Configuration
                </Typography.Title>
              </Col>
              {renderAmazonS3FormItems()}
              <Col offset={6} style={{ marginBottom: "10px" }}>
                <Typography.Link
                  onClick={() => {
                    setShowFilter(!showFilter);
                  }}
                >
                  <CaretRightOutlined rotate={showFilter ? 90 : 0} />
                  Advanced filter
                </Typography.Link>
              </Col>
              {showFilter && (
                <>
                  <Form.Item
                    label="Max object size"
                    key="max_obj_size"
                    name="max_obj_size"
                  >
                    <InputNumber />
                  </Form.Item>
                  <Form.Item
                    label="Creation Date"
                    key="creation_date"
                    name="creation_date"
                  >
                    <DatePicker allowClear={false} />
                  </Form.Item>
                </>
              )}
              <Col offset={6}>
                <Typography.Title level={4}>
                  Lyve S3 Configuration
                </Typography.Title>
              </Col>
              {renderLyveS3FormItems()}
              <Form.Item wrapperCol={{ offset: 6 }}>
                <Button type="primary" htmlType="submit">
                  Submit
                </Button>
              </Form.Item>
            </Form>
          </Col>
        </Row>
      </Card>
    </div>
  );
};

export default CreateMigration;
