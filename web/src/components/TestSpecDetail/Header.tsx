import {ArrowLeftOutlined} from '@ant-design/icons';
import {EditorView} from '@codemirror/view';
import {Button, Divider} from 'antd';

import useEditorTheme from 'components/AdvancedEditor/hooks/useEditorTheme';
import TestSpecActions from 'components/TestSpec/Actions';
import * as STestSpec from 'components/TestSpec/TestSpec.styled';
import {singularOrPlural} from 'utils/Common';
import {tracetest} from 'utils/grammar';
import * as S from './TestSpecDetail.styled';

interface IProps {
  affectedSpans: number;
  assertionsFailed: number;
  assertionsPassed: number;
  isDeleted: boolean;
  isDraft: boolean;
  onClose(): void;
  onDelete(): void;
  onEdit(): void;
  onRevert(): void;
  title: string;
}

const Header = ({
  affectedSpans,
  assertionsFailed,
  assertionsPassed,
  isDeleted,
  isDraft,
  onClose,
  onDelete,
  onEdit,
  onRevert,
  title,
}: IProps) => {
  const editorTheme = useEditorTheme({isEditable: false});

  return (
    <>
      <S.HeaderContainer>
        <S.Row $hasGap>
          <Button icon={<ArrowLeftOutlined />} onClick={onClose} type="link" />
          <S.HeaderTitle level={2}>Test Spec Detail</S.HeaderTitle>
          <div>
            {Boolean(assertionsPassed) && (
              <STestSpec.HeaderDetail>
                <STestSpec.HeaderDot $passed />
                {assertionsPassed}
              </STestSpec.HeaderDetail>
            )}
            {Boolean(assertionsFailed) && (
              <STestSpec.HeaderDetail>
                <STestSpec.HeaderDot $passed={false} />
                {assertionsFailed}
              </STestSpec.HeaderDetail>
            )}
            <STestSpec.HeaderDetail>
              <STestSpec.HeaderSpansIcon />
              {`${affectedSpans} ${singularOrPlural('span', affectedSpans)}`}
            </STestSpec.HeaderDetail>
          </div>
        </S.Row>

        <TestSpecActions
          isDeleted={isDeleted}
          isDraft={isDraft}
          onDelete={onDelete}
          onEdit={onEdit}
          onRevert={onRevert}
        />
      </S.HeaderContainer>

      <Divider />

      <STestSpec.HeaderTitle
        data-cy="advanced-selector"
        editable={false}
        extensions={[tracetest(), editorTheme, EditorView.lineWrapping]}
        maxHeight="120px"
        placeholder="Selecting All Spans"
        spellCheck={false}
        value={title || 'All Spans'}
      />

      <Divider />
    </>
  );
};

export default Header;