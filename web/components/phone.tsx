import { Descriptions } from "antd";
import { Phone } from "../protobuf/phone/phone_pb";

interface PhoneComponentProp {
  phone: Phone.AsObject;
}

const PhoneComponent = ({ phone }: PhoneComponentProp): JSX.Element => (
  <div>
    <h1>{phone.name}</h1>
    <Descriptions
      bordered
      column={{
        xxl: 4,
        xl: 3,
        lg: 3,
        md: 3,
        sm: 2,
        xs: 1,
      }}
    >
      <Descriptions.Item label="Name">{phone.name}</Descriptions.Item>
      <Descriptions.Item label="Make">
        <a href={`/make/${phone.make.id}`}>{phone.make.name}</a>
      </Descriptions.Item>
      <Descriptions.Item label="OS">
        <a href={`/os/${phone.os.id}`}>{phone.os.name}</a>
      </Descriptions.Item>
    </Descriptions>
  </div>
);

export default PhoneComponent;
