import {Form} from 'antd';
import {useCallback, useEffect, useState} from 'react';
import {HTTP_METHOD} from 'constants/Common.constants';
import {useCreateTest} from 'providers/CreateTest/CreateTest.provider';
import CreateStepFooter from 'components/CreateTestSteps/CreateTestStepFooter';
import * as Step from 'components/CreateTestPlugins/Step.styled';
import {TRequestAuth, THTTPRequest} from 'types/Test.types';
import RequestDetailsForm from './RequestDetailsForm';

export interface IRequestDetailsValues {
  body: string;
  auth: TRequestAuth;
  headers: THTTPRequest['headers'];
  method: HTTP_METHOD;
  name: string;
  url: string;
}

const RequestDetails = () => {
  const [isFormValid, setIsFormValid] = useState(false);
  const [form] = Form.useForm<IRequestDetailsValues>();
  const {onNext} = useCreateTest();
  const {
    draftTest: {serviceUnderTest: {triggerSettings: {http: request = {}} = {}} = {}},
  } = useCreateTest();

  const handleNext = useCallback(() => {
    form.submit();
  }, [form]);

  const handleSubmit = useCallback(
    (values: IRequestDetailsValues) => {
      onNext({serviceUnderTest: {triggerSettings: {http: values}}});
    },
    [onNext]
  );

  const onRefreshData = useCallback(async () => {
    const {url = '', body = '', method = HTTP_METHOD.GET} = request;

    form.setFieldsValue({url, body, method: method as HTTP_METHOD});

    try {
      await form.validateFields();
      setIsFormValid(true);
    } catch (err) {
      setIsFormValid(false);
    }
  }, [form, request]);

  useEffect(() => {
    onRefreshData();
  }, [onRefreshData]);

  return (
    <Step.Step>
      <Step.FormContainer>
        <Step.Title>Provide additional information</Step.Title>
        <RequestDetailsForm form={form} onSubmit={handleSubmit} onValidation={setIsFormValid} />
      </Step.FormContainer>
      <CreateStepFooter isValid={isFormValid} onNext={handleNext} />
    </Step.Step>
  );
};

export default RequestDetails;