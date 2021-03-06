import React from "react";
import { StatusCode } from "grpc-web";
import Container from "../../components/container";
import SEO from "../../components/seo";
import Custom404 from "../404";
import { GetOneByIDRequest } from "../../protobuf/phone/phone_reader_pb";
import { Phone } from "../../protobuf/phone/phone_pb";
import PhoneReaderClient from "../../clients/nodejs/phone_reader_client";
import PhoneComponent from "../../components/phone";
import {
  GetOneByIDServerSide,
  GetServerSideFunc,
  GetOnePageProp,
} from "../../components/getOne";

export const getServerSideProps: GetServerSideFunc<Phone.AsObject> = GetOneByIDServerSide<
  Phone.AsObject,
  Phone
>(new GetOneByIDRequest(), PhoneReaderClient, "legit");

const PhonePage = (props: GetOnePageProp<Phone.AsObject>): JSX.Element => {
  if (props.result) {
    return (
      <Container defKey="1">
        <SEO title={props.result.name} />
        <PhoneComponent phone={props.result} />
      </Container>
    );
  }

  if (props.error.code === StatusCode.NOT_FOUND) {
    return (
      <Custom404
        defKey="1"
        title="Phone not found"
        message={`Phone with id ${props.id} not found.`}
      />
    );
  }

  return (
    <Container defKey="1">
      <h1>{props.httpStatusCode}</h1>
      <p>{props.error.message}</p>
    </Container>
  );
};

export default PhonePage;
