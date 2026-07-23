import { createSlice } from "@reduxjs/toolkit";
import type { GroupSchema } from "../types/types";
import { getUserGroups } from "../api/getUsergroups";
import { getGroupsByTutorId } from "../api/getGroupsByTutorId";
import { createGroup } from "../api/createGroup";
import { deleteGroup } from "../api/deleteGroup";
import { removeUserFromGroup } from "../api/removeUserFromGroup";
import { updateGroup } from "../api/updateGroup";

const initialState: GroupSchema = {
    data: null,
    isLoading: false,
    error: undefined
};

const groupSlice = createSlice({
    name: 'group',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(getUserGroups.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getUserGroups.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getUserGroups.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = action.payload.map(g => {
                return {
                    id: g.id,
                    title: g.title,
                    description: g.description,
                    subject: g.subject,
                    users: g.users?.map(gu => {
                        return {
                            id: gu.id,
                            name: gu.name,
                            surname: gu.surname,
                            patronymic: gu.patronymic,
                            tgUsername: gu.tg_username,
                        }
                    }),
                    tutorId: g.tutor_id,
                    tgChatId: g.tg_chat_id,
                    tgGroupLink: g.tg_group_link,
                }
            })
        })
        builder.addCase(getGroupsByTutorId.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getGroupsByTutorId.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getGroupsByTutorId.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = action.payload.map(g => {
                return {
                    id: g.id,
                    title: g.title,
                    description: g.description,
                    subject: g.subject,
                    users: g.users?.map(gu => {
                        return {
                            id: gu.id,
                            name: gu.name,
                            surname: gu.surname,
                            patronymic: gu.patronymic,
                            tgUsername: gu.tg_username,
                        }
                    }),
                    tutorId: g.tutor_id,
                    tgChatId: g.tg_chat_id,
                    tgGroupLink: g.tg_group_link,
                }
            })
        })
        builder.addCase(createGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(createGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(createGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data?.push({
                id: action.payload.id,
                title: action.payload.title,
                description: action.payload.description,
                users: action.payload.users,
                subject: action.payload.subject,
                tutorId: action.payload.tutor_id,
                tgGroupLink: action.payload.tg_group_link
            })
        })
        builder.addCase(deleteGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(deleteGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = state.data?.filter(g => g.id != action.meta.arg) || null
        })

        builder.addCase(removeUserFromGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(removeUserFromGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(removeUserFromGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            const group = state.data?.find(
                g => g.id === action.meta.arg.groupID
            );

            if (group && group.users) {
                group.users = group.users.filter(
                    u => u.id !== action.meta.arg.userID
                );
            }
        })
        builder.addCase(updateGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(updateGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateGroup.fulfilled, (state, action) => {
            state.isLoading = false;
            state.error = '';

            const group = state.data?.find(g => g.id === action.meta.arg.groupID);

            console.log(action.payload)

            if (group) {
                group.title = action.payload.title;
                group.description = action.payload.description;
                group.tutorId = action.payload.tutor_id;
                group.tgGroupLink = action.payload.tg_group_link;
            }
        });
    }
});

export const { actions: groupActions, reducer: groupReducer } =
    groupSlice;