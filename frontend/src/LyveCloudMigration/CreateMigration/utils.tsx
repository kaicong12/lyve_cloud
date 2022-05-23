import React from 'react';
import { Form, Input } from 'antd';

export const renderAmazonS3FormItems = () => {
  const nameLists: {
    value: string;
    label: string;
  }[] = [
    {
      value: 'accessKey',
      label: 'Access key',
    },
    {
      value: 'secretAccessKey',
      label: 'Secret access key',
    },
    {
      value: 'bucketName',
      label: 'Bucket name',
    },
    {
      value: 'path',
      label: 'Path',
    },
  ];
  return nameLists.map((ele) => {
    return (
      <Form.Item
        label={ele.label}
        key={ele.value}
        name={ele.value}
        rules={[{ required: true, message: `Please enter the ${ele.label} !` }]}
      >
        <Input></Input>
      </Form.Item>
    );
  });
};

export const renderLyveS3FormItems = () => {
  const nameLists: {
    value: string;
    label: string;
  }[] = [
    {
      value: 'lyveAccessKey',
      label: 'Access key',
    },
    {
      value: 'lyveSecretAccessKey',
      label: 'Secret access key',
    },
    {
      value: 'lyveBucketName',
      label: 'Bucket name',
    },
  ];
  return nameLists.map((ele) => {
    return (
      <Form.Item
        label={ele.label}
        key={ele.value}
        name={ele.value}
        rules={[{ required: true, message: `Please enter the Lyve ${ele.label} !` }]}
      >
        <Input></Input>
      </Form.Item>
    );
  });
};
